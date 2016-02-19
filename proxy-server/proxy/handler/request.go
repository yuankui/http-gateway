package handler

import (
	"errors"
	"fmt"
	"http-gateway/models"
	"http-gateway/proxy-server/cache"
	"http-gateway/proxy-server/proxy"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func RequestHandler(request *proxy.ProxyRequest, rw http.ResponseWriter) (ret bool) {

	err := parseRequest(request)

	if err != nil {
		log.Println(err)
		rw.WriteHeader(400)
		rw.Write([]byte("no site found in url"))
		return true
	}

	s, found := cache.SiteCache.Get(request.SiteName)

	if !found {
		log.Println("domain:", request.SiteName, " not found")
		rw.WriteHeader(404)
		rw.Write([]byte("site not found"))

		startTime := request.Context["start"].(time.Time)
		cost := (time.Now().UnixNano() - startTime.UnixNano()) / 1000000

		log.Printf("[%d] %d %s -> %s%s", 404, cost, request.SiteName, request.NewRequest.Host, request.NewRequest.RequestURI)
		return true
	}

	site := s.(*models.Site)

	newUrl := fmt.Sprintf("http://%s:%d%s", site.Host, site.Port, request.NewRequest.RequestURI)
	u, err := url.Parse(newUrl)

	if err != nil {
		fmt.Println(err)
	}

	request.NewRequest.Host = fmt.Sprintf("%s:%d", site.Host, site.Port)
	request.NewRequest.URL = u
	return false
}

func parseRequest(req *proxy.ProxyRequest) error {
	split := strings.SplitN(req.OldRequest.Host, ".", 2)
	if len(split) != 2 {
		return errors.New("domain not found in host")
	}

	req.SiteName = split[0]
	return nil
}
