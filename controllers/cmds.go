package controllers

import (
	"strings"
	"webjokes/cmds"
	"webjokes/models"

	"github.com/astaxie/beego"
)

const (
	ADMIN_USER string = "ADMIN_USER"
)

type CmdsControllers struct {
	beego.Controller
}

func (this *CmdsControllers) Get() {
	arrCmd := this.operateCmd()
	if arrCmd[0] == "" {
		this.OutJosn("Error", "命令参数无效")
		return
	}
	il := len(arrCmd)
	// 判断是否有登入
	if !this.CheckLogin() {
		if arrCmd[0] == "login" {
			// 执行登入
			if il < 2 {
				this.OutJosn("LOGIN", map[string]interface{}{
					"result": false,
					"info":   "请输入正确登入码",
				})
				return
			}
			// 判断密码是否正确
			if ptAdmin, err := models.CheckAdminUser(arrCmd[1]); err != nil {
				this.OutJosn("LOGIN", map[string]interface{}{
					"result": false,
					"info":   "登入码不正确",
				})
			} else {
				// 写入SESSION
				this.SetSession(ADMIN_USER, ptAdmin)
				// 返回
				this.OutJosn("LOGIN", map[string]interface{}{
					"result": true,
					"info":   "ok",
				})
			}
			return
		}
		this.OutJosn("LoginOut", "is not login")
		return
	}

	// 转由命令处理对象处理
	res, err := cmds.RunCmd(this, arrCmd[0], arrCmd)
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

func (this *CmdsControllers) OutJosn(cmd string, info interface{}) {
	this.Data["json"] = models.CreateJsonDb(cmd, info)
	this.ServeJSON()
}
