package proxy

import (
	"bytes"
	"net/http"
)

type Response struct {
	Header     http.Header
	StatusCode int
	bodyBuffer bytes.Buffer
	Cookies    []*http.Cookie
}

func (this *Response) Write(p []byte) (n int, err error) {
	return this.bodyBuffer.Write(p)
}

func (this *Response) Body() []byte {
	return this.bodyBuffer.Bytes()
}

func NewResponse() Response {
	res := Response{Header: make(http.Header)}
	return res
}
