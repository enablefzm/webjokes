package cmds

import (
	"vava6/vatools"
	"webjokes/models"
)

// 随机获取当前命令对象
func joke(ptController IFController, cmd string, parames []string) *CmdResult {
	// return nil, fmt.Errorf("error")
	il := len(parames)
	if il < 1 {
		return createResult("ERROR", false, "缺少参数")
	}
	switch parames[0] {
	// 随机获取对象
	case "rnd":
		ptJoke, ok := models.OBJokePool.Get()
		if !ok {
			return createResult("JOKE_RND", false, "没有需要过审的笑话")
		}
		return createResult("JOKE_RND", true, ptJoke.GetInfo())
	// 审核
	// joke check <id> <state 1|2>
	case "check":
		if il < 3 {
			return createResult("ERROR", false, "你要审核哪个笑话")
		}
		jokeID := vatools.SInt(parames[1])
		state := int8(vatools.SInt(parames[2]))
		ptJoke, err := models.NewJokeSourceOnID(jokeID)
		if err != nil {
			return createResult("JOKE_CHECK", false, err.Error())
		}
		// 获得修改对象
		ptAdmin, ok := ptController.GetAdminUser()
		if !ok {
			return createError("没有找到相应的用户对象")
		}
		// 获取到笑话指定修改者ID
		ptJoke.SetCheckIDs(ptAdmin.GetID())
		ptJoke.SetCheckState(state)
		if il > 3 {
			lbls := parames[3]
			ptJoke.SetLabels(lbls)
		}
		ptJoke.Updata()
		// 更新Admin的操作数量
		ptAdmin.AddCheck()
		return createResult("JOKE_CHECK", true, "")
	}
	return createResult("ERROR", false, "没有相应的参数")
}

func init() {
	RegCmd("joke", joke)
}
