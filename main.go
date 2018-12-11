package main

import (
	_ "webjokes/routers"

	"github.com/astaxie/beego"
)

func main() {
	// beego.SetLevel(beego.LevelError)
	// beego.BeeLogger.DelLogger("console")
	// beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	// beego.BConfig.Log.AccessLogs = false
	beego.Run()
}
