package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	uuid := this.GetString("uuid", "jimmy")
	this.Ctx.WriteString(fmt.Sprintf("<h1>我叫%s</h1>", uuid))
}
