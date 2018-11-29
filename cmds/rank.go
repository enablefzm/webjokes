package cmds

import (
	"webjokes/models"
)

func rank(ptController IFController, cmd string, parames []string) *CmdResult {
	rss, err := models.DBSave.QuerysLimit("name,checks,logins", "admin_users", "checks > 0", 1, 100, "checks DESC")
	if err != nil {
		return createError(err.Error())
	}
	return createResult("RANK", true, rss)
}

func init() {
	RegCmd("rank", rank)
}
