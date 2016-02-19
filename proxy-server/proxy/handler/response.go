package handler

import (
	"http-gateway/proxy-server/proxy"
	"log"
	"time"
)

func ResponseHandler(request *proxy.ProxyRequest, response *proxy.Response) *proxy.Response {
	//fmt.Println(string(response.Body()))

	for _, cookie := range response.Cookies {
		// remove domain
		cookie.Domain = ""
	}

	startTime := request.Context["start"].(time.Time)
	cost := (time.Now().UnixNano() - startTime.UnixNano()) / 1000000

	log.Printf("[%d] %d %s -> %s%s", response.StatusCode, cost, request.SiteName, request.NewRequest.Host, request.NewRequest.RequestURI)
	return response
}
