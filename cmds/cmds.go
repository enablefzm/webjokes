package cmds

import (
	"fmt"
	"webjokes/models"
)

type IFController interface {
	GetAdminUser() (*models.AdminUser, bool)
}

type CmdResult struct {
	CmdKey   string
	CmdValue interface{}
}

type docmd func(ptController IFController, parames []string) (*CmdResult, error)

var mpCmds map[string]docmd = make(map[string]docmd, 10)

// 执行命令对象
func RunCmd(ptController IFController, cmd string, parames []string) (*CmdResult, error) {
	// 判断当前是否有这个命令存在
	if fnCmd, ok := mpCmds[cmd]; !ok {
		return nil, fmt.Errorf("命令不存在")
	} else {
		return fnCmd(ptController, parames)
	}
}

func RegCmd(cmdKey string, fnCmd docmd) {
	mpCmds[cmdKey] = fnCmd
}
