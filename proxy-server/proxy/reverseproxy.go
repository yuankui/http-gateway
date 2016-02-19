// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP reverse proxy handler

package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// onExitFlushLoop is a callback set by tests to detect the state of the
// flushLoop() goroutine.
var onExitFlushLoop func()

// ReverseProxy is an HTTP Handler that takes an incoming request and
// sends it to another server, proxying the response back to the
// client.
type ReverseProxy struct {
	// Director must be a function which modifies
	// the request into a new request to be sent
	// using Transport. Its response is then copied
	// back to the original client unmodified.
	RequestHandler func(*ProxyRequest, http.ResponseWriter) (ret bool)

	// The transport used to perform proxy requests.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper

	// FlushInterval specifies the flush interval
	// to flush to the client while copying the
	// response body.
	// If zero, no periodic flushing is done.
	FlushInterval time.Duration

	// ErrorLog specifies an optional logger for errors
	// that occur when attempting to proxy the request.
	// If nil, logging goes to os.Stderr via the log package's
	// standard logger.
	ErrorLog *log.Logger

	ResponseHandler func(*ProxyRequest, *Response) *Response
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

type requestCanceler interface {
	CancelRequest(*http.Request)
}

type runOnFirstRead struct {
	io.Reader // optional; nil means empty body

	fn func() // Run before first Read, then set to nil
}

func (c *runOnFirstRead) Read(bs []byte) (int, error) {
	if c.fn != nil {
		c.fn()
		c.fn = nil
	}
	if c.Reader == nil {
		return 0, io.EOF
	}
	return c.Reader.Read(bs)
}

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	transport := p.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	// outreq := new(http.Request)
	proxyRequest := ProxyRequest{OldRequest: req}
	proxyRequest.Context = make(map[string]interface{})
	proxyRequest.Context["start"] = time.Now()
	proxyRequest.NewRequest = new(http.Request)
	*proxyRequest.NewRequest = *req // includes shallow copies of maps, but okay

	if closeNotifier, ok := rw.(http.CloseNotifier); ok {
		if requestCanceler, ok := transport.(requestCanceler); ok {
			reqDone := make(chan struct{})
			defer close(reqDone)

			clientGone := closeNotifier.CloseNotify()

			proxyRequest.NewRequest.Body = struct {
				io.Reader
				io.Closer
			}{
				Reader: &runOnFirstRead{
					Reader: proxyRequest.NewRequest.Body,
					fn: func() {
						go func() {
							select {
							case <-clientGone:
								requestCanceler.CancelRequest(proxyRequest.NewRequest)
							case <-reqDone:
							}
						}()
					},
				},
				Closer: proxyRequest.NewRequest.Body,
			}
		}
	}

	ret := p.RequestHandler(&proxyRequest, rw)
	if ret {
		return
	}
	proxyRequest.NewRequest.Proto = "HTTP/1.1"
	proxyRequest.NewRequest.ProtoMajor = 1
	proxyRequest.NewRequest.ProtoMinor = 1
	proxyRequest.NewRequest.Close = false

	// Remove hop-by-hop headers to the backend.  Especially
	// important is "Connection" because we want a persistent
	// connection, regardless of what the client sent to us.  This
	// is modifying the same underlying map from req (shallow
	// copied above) so we only copy it if necessary.
	copiedHeaders := false
	for _, h := range hopHeaders {
		if proxyRequest.NewRequest.Header.Get(h) != "" {
			if !copiedHeaders {
				proxyRequest.NewRequest.Header = make(http.Header)
				copyHeader(proxyRequest.NewRequest.Header, req.Header)
				copiedHeaders = true
			}
			proxyRequest.NewRequest.Header.Del(h)
		}
	}

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		// If we aren't the first proxy retain prior
		// X-Forwarded-For information as a comma+space
		// separated list and fold multiple headers into one.
		if prior, ok := proxyRequest.NewRequest.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		proxyRequest.NewRequest.Header.Set("X-Forwarded-For", clientIP)
	}

	res, err := transport.RoundTrip(proxyRequest.NewRequest)

	originResponse := NewResponse()

	if err != nil {
		p.logf("http: proxy error: %v", err)
		originResponse.StatusCode = http.StatusAccepted
	} else {
		p.copyResponse(&originResponse, res)
	}

	var changedResponse *Response
	if p.ResponseHandler != nil {

		if err != nil {
			rw.WriteHeader(400)
			rw.Write([]byte(err.Error()))
			return
		}
		changedResponse = p.ResponseHandler(&proxyRequest, &originResponse)
	}

	p.WriteResponse(changedResponse, rw)
}

func (this *ReverseProxy) WriteResponse(response *Response, rw http.ResponseWriter) {
	copyHeader(rw.Header(), response.Header)

	for _, cookie := range response.Cookies {
		rw.Header().Add("Set-Cookie", cookie.String())
	}

	rw.WriteHeader(response.StatusCode)
	rw.Write(response.Body())
}

func (p *ReverseProxy) copyResponse(originResponse *Response, res *http.Response) {
	// 拷贝body
	io.Copy(originResponse, res.Body)
	res.Body.Close() // close now, instead of defer, to populate res.Trailer

	// 拷贝statusCode
	originResponse.StatusCode = res.StatusCode

	// 拷贝header
	for _, h := range hopHeaders {
		res.Header.Del(h)
	}

	copyHeader(originResponse.Header, res.Header)
	copyHeader(originResponse.Header, res.Trailer)

	// 设置cookie引用
	originResponse.Cookies = res.Cookies()
}

func (p *ReverseProxy) logf(format string, args ...interface{}) {
	if p.ErrorLog != nil {
		p.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

type writeFlusher interface {
	io.Writer
	http.Flusher
}

type maxLatencyWriter struct {
	dst     writeFlusher
	latency time.Duration

	lk   sync.Mutex // protects Write + Flush
	done chan bool
}
