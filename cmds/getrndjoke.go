package cmds

import (
	"fmt"
)

// 随机获取当前命令对象
func GetRndJoke(ptController IFController, cmd string, parames []string) (*CmdResult, error) {
	return nil, fmt.Errorf("error")
}

func init() {
	RegCmd("getRndJoke", GetRndJoke)
}
