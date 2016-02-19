package cache

import (
	"testing"
	"log"
)

func TestCache(t *testing.T) {
	site, found := SiteCache.Get("domain1")

	log.Println(site, found)
}
