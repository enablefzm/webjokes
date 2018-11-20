package routers

import (
	"webjokes/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/cmds", &controllers.CmdsControllers{})
}
