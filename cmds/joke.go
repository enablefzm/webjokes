package cmds

import (
	"fmt"
)

// 随机获取当前命令对象
func joke(ptController IFController, cmd string, parames []string) (*CmdResult, error) {
	return nil, fmt.Errorf("error")
}

func init() {
	RegCmd("joke", joke)
}
