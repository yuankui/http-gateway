package cache

import (
	"time"
	"github.com/astaxie/beego/orm"
	"http-gateway/models"
	"log"
)

var SiteCache *LoadCache
var o orm.Ormer

func loadSite(name string) (interface{}, bool) {
	site := models.Site{Domain:name}
	error := o.Read(&site)

	log.Println("loading site:", site)
	if error != nil {
		log.Println(error)
		return nil, false
	}

	return &site, true
}

func init() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 30 seconds

	SiteCache = NewCache(time.Hour * 24 * 30, time.Second * 10, loadSite)

	o = orm.NewOrm()

	go func() {
		var sites []models.Site

		o.QueryTable("site").All(&sites)

		for _, site := range sites {
			copy := new(models.Site)
			*copy = site
			SiteCache.Set(site.Domain, copy)
		}
	}()
}
