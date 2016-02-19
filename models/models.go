package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Site struct {
	Domain      string    `orm:"pk" form:"domain"`
	Host        string    `form:"host"`
	Port        int       `form:"port"`
	DefaultPath string    `form:"defaultPath"`
	Description string    `form:"description"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Site))
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("connectionString"), 30)

	orm.RunSyncdb("default", false, true)
}
