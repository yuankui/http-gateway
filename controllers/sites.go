package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"http-gateway/models"
	"reflect"
	"http-gateway/logger"
	"github.com/davecgh/go-spew/spew"
)

type SitesController struct {
	beego.Controller
}

func (c *SitesController) common() {
	c.Data["SiteDomain"] = beego.AppConfig.String("domain")
}

func (c *SitesController) Index() {
	c.common()
	o := orm.NewOrm()

	var sites []models.Site

	num, _ := o.QueryTable("site").All(&sites)
	logger.Logger.Info("Returned Rows Num: %d\n", num)

	c.Data["Sites"] = sites
	c.Layout = "layout.tpl"
	c.TplNames = "sites/index.tpl"
}

func (c *SitesController) Edit() {
	c.common()
	o := orm.NewOrm()

	domain := c.Ctx.Input.Param(":domain")

	site := models.Site{Domain: domain}

	err := o.Read(&site)

	if err != nil {
		beego.Warn(reflect.TypeOf(err), err)
		c.Abort("404")
	}

	c.Data["Site"] = site
	c.Data["Disabled"] = "disabled"
	c.Layout = "layout.tpl"
	c.TplNames = "sites/edit.tpl"
}

func (c *SitesController) RedirectHome() {
	c.Redirect("/sites", 301)
}

func (c*SitesController) New() {
	c.common()
	c.TplNames = "sites/edit.tpl"
	c.Layout = "layout.tpl"
}

func (this*SitesController) Update() {
	site := models.Site{}
	this.ParseForm(&site)
	site.Domain = this.Ctx.Input.Param(":domain")

	o := orm.NewOrm()

	oldSite := models.Site{Domain:site.Domain}

	spew.Dump(oldSite)
	error := o.Read(&oldSite)

	if error != nil {
		logger.Logger.Error(error.Error())
		this.Abort("500")
	}

	_, error = o.Update(&site)
	if error != nil {
		logger.Logger.Error(error.Error())
		this.Abort("500")
		return
	}
	this.RedirectHome()
}

func (this*SitesController) Create() {
	site := models.Site{}
	this.ParseForm(&site)

	o := orm.NewOrm()

	oldSite := models.Site{Domain:site.Domain}

	// 检查是否已经存在
	error := o.Read(&oldSite)
	if error == nil {
		logger.Logger.Error(error.Error())
		this.Abort("500")
	}

	// 插入
	_, error = o.Insert(&site)
	if error != nil {
		logger.Logger.Error(error.Error())
		this.Abort("500")
		return
	}
	this.RedirectHome()
}

func (this*SitesController) Delete() {
	domain := this.Ctx.Input.Param(":domain")

	o := orm.NewOrm()

	_, error := o.Delete(&models.Site{Domain:domain})

	if error != nil {
		logger.Logger.Error(error.Error())
		this.Abort("500")
	}

	this.RedirectHome()
}