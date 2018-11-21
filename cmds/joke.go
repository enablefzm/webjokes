package cmds

import (
	"fmt"
	"webjokes/models"
)

// 随机获取当前命令对象
func joke(ptController IFController, cmd string, parames []string) (*CmdResult, error) {
	// return nil, fmt.Errorf("error")
	il := len(parames)
	if il < 1 {
		return nil, createResult("ERROR", false, "缺少参数")
	}
	switch parames[0] {
	// 随机获取对象
	case "rnd":
		// models.DBSave.Querys()
		// models.OBJokePool.Get()

	}
}

func init() {
	RegCmd("joke", joke)
}
