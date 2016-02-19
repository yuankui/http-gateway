package logger

import (
	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

func init() {
	Logger = logs.NewLogger(10000)

	Logger.SetLogger("console", "")

	Logger.EnableFuncCallDepth(true)
}
