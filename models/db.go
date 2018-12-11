package models

import (
	"fmt"
	"os"
	"vava6/mysql"
)

var DBSave *mysql.DBs

func init() {
	err := LinkDBServer()
	if err != nil {
		os.Exit(1)
	}
}

func LinkDBServer() error {
	var err error

	// 加载配置文件
	cfg := mysql.NewCfg()
	cfg.DBName = "joke"
	cfg.Address = "119.23.235.220"
	cfg.User = "joke"
	cfg.Pass = "jokeUser&2018"
	cfg.MaxConn = 50

	fmt.Print("【正在连接数据库...")
	DBSave, err = mysql.NewDBs(
		cfg.DBName,
		cfg.Address,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.MaxConn,
		cfg.MinConn,
	)
	if err != nil {
		fmt.Print("连接数据库失败-", err.Error())
	} else {
		fmt.Print("连接数据库成功")
	}
	fmt.Println("】")
	return err
}
