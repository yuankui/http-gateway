package proxy

import (
	"net/http"
)

type ProxyRequest struct {
	OldRequest *http.Request
	NewRequest *http.Request
	SiteName   string
	Context    map[string]interface{}
}
