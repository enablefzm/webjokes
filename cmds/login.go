package cmds

import (
	"fmt"
	"webjokes/models"

	"github.com/astaxie/beego/context"
)

func login(ptController IFController, cmd string, parames []string) *CmdResult {
	// 判断是否已经登入游戏
	if ptAdmin, ok := ptController.GetAdminUser(); ok {
		return createResult("LOGIN", false, fmt.Sprint(ptAdmin.GetUid(), "已经登入系统"))
	}
	il := len(parames)
	if il < 1 {
		return createResult("LOGIN", false, "请输入正确的登入码")
	}
	// 判断密码是否正确
	ptAdmin, err := models.CheckAdminUser(parames[0])
	if err != nil {
		return createResult("LOGIN", false, "登入码不正确")
	}
	// 设定Admin
	ptController.SetAdminUser(ptAdmin)
	ptController.GetCtx()
	// 执行LOG信息
	models.DBSave.Insert("login_logs", map[string]interface{}{
		"uid":      ptAdmin.GetUid(),
		"agent":    ptController.GetCtx().Request.UserAgent(),
		"remoteIp": getRemoteIp(ptController.GetCtx()),
	})
	// 更新最后一次登入时间和登入次数
	ptAdmin.AddLogin()
	// 返回登入成功
	return createResult("LOGIN", true, "登入成功")
}

func getRemoteIp(pt *context.Context) string {
	ip := pt.Request.Header.Get("X-Forwarded-For")
	fmt.Println("[RemoteIP] X-Forwarded-For ", ip, " | RemoteAddr ", pt.Request.RemoteAddr)
	if len(ip) > 5 {
		return ip
	}
	return pt.Request.RemoteAddr
}

func init() {
	RegCmd("login", login)
}
