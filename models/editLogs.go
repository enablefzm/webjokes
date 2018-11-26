package models

import (
	"database/sql"
)

func LogEditContent(jokeID int, jokeContent string, userID int, userUid string) (sql.Result, error) {
	return DBSave.Insert("edit_logs", map[string]interface{}{
		"jokeID":      jokeID,
		"jokeContent": jokeContent,
		"editUserID":  userID,
		"editUserUid": userUid,
	})
}
