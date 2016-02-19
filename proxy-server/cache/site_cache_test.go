package cache

import (
	"testing"
	"fmt"
	"time"
)

func TestIt(t *testing.T) {
	fmt.Println(SiteCache.Get("domain1"))


	fmt.Println(SiteCache.Get("domain1"))
	time.Sleep(time.Second * 3)

	fmt.Println(SiteCache.Get("domain1"))

	time.Sleep(10 * time.Second)

	fmt.Println(SiteCache.Get("domain1"))

	time.Sleep(3 * time.Second)
}