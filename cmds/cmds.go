package cmds

import (
	"fmt"
	"webjokes/models"

	"github.com/astaxie/beego/context"
)

type IFController interface {
	GetAdminUser() (*models.AdminUser, bool)
	SetAdminUser(ptAdmin *models.AdminUser)
	SetSession(name interface{}, value interface{})
	GetCtx() *context.Context
}

type CmdResult struct {
	CmdKey   string
	CmdValue interface{}
}

type docmd func(ptController IFController, cmd string, parames []string) (*CmdResult, error)

var mpCmds map[string]docmd = make(map[string]docmd, 10)

// 执行命令对象
func RunCmd(ptController IFController, cmd string, parames []string) (*CmdResult, error) {
	// 判断当前是否有这个命令存在
	if fnCmd, ok := mpCmds[cmd]; !ok {
		return nil, fmt.Errorf("命令不存在")
	} else {
		return fnCmd(ptController, cmd, parames)
	}
}

func RegCmd(cmdKey string, fnCmd docmd) {
	mpCmds[cmdKey] = fnCmd
}

func createResult(key string, success bool, msg string) *CmdResult {
	return &CmdResult{
		CmdKey: key,
		CmdValue: map[string]interface{}{
			"result": success,
			"info":   msg,
		},
	}
}
