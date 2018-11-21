package controllers

import (
	"strings"
	"webjokes/cmds"
	"webjokes/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	ADMIN_USER string = "ADMIN_USER"
)

type CmdsControllers struct {
	beego.Controller
}

func (this *CmdsControllers) Get() {
	this.doing()
}

func (this *CmdsControllers) Post() {
	this.ServeJSON()
	// this.doing()
}

func (this *CmdsControllers) doing() {
	//	fmt.Println(this.Ctx.Request.RemoteAddr, this.Ctx.Request.UserAgent())
	//	fmt.Println(this.GetString("cmd"))
	arrCmd := this.operateCmd()
	if arrCmd[0] == "" {
		this.OutJosn("Error", "命令参数无效")
		return
	}
	// 命令
	cmd := arrCmd[0]
	// 参数
	parames := arrCmd[1:]

	// 判断是否有登入
	if !this.CheckLogin() && cmd != "login" {
		this.OutJosn("LoginOut", "is not login")
		return
	}

	// 转由命令处理对象处理
	res, err := cmds.RunCmd(this, cmd, parames)
	if err != nil {
		this.OutJosn("ERROR", err.Error())
	}
	this.OutJosn(res.CmdKey, res.CmdValue)
}

func (this *CmdsControllers) operateCmd() []string {
	cmd := strings.Trim(this.GetString("cmd"), " ")
	arr := strings.Split(cmd, " ")
	return arr
}

func (this *CmdsControllers) CheckLogin() bool {
	if this.GetSession(ADMIN_USER) == nil {
		return false
	}
	return true
}

func (this *CmdsControllers) GetCtx() *context.Context {
	return this.Ctx
}

func (this *CmdsControllers) GetAdminUser() (*models.AdminUser, bool) {
	v := this.GetSession(ADMIN_USER)
	if v == nil {
		return nil, false
	}
	ptAdmin, ok := v.(*models.AdminUser)
	if !ok {
		return nil, false
	}
	return ptAdmin, true
}

func (this *CmdsControllers) SetAdminUser(ptAdmin *models.AdminUser) {
	this.SetSession(ADMIN_USER, ptAdmin)
}

func (this *CmdsControllers) OutJosn(cmd string, info interface{}) {
	this.Data["json"] = models.CreateJsonDb(cmd, info)
	this.ServeJSON()
}
