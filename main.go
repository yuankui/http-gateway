package main

import (
	"github.com/astaxie/beego"
	_ "http-gateway/models"
	_ "http-gateway/routers"
)

func main() {
	beego.Run()
}
