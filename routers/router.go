package routers

import (
	"http-gateway/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/sites", &controllers.SitesController{}, "get:Index")
	beego.Router("/sites/new", &controllers.SitesController{}, "get:New")
	beego.Router("/sites/new", &controllers.SitesController{}, "post:Create")
	beego.Router("/sites/:domain/", &controllers.SitesController{}, "post:Update")
	beego.Router("/sites/:domain/", &controllers.SitesController{}, "get:Edit" )
	beego.Router("/sites/:domain/delete", &controllers.SitesController{}, "post:Delete" )
	beego.Router("/", &controllers.SitesController{}, "get:RedirectHome")

	beego.ErrorController(&controllers.ErrorController{})
	beego.StaticDir["/static"] = "static"
}
