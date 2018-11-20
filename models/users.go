package models

import (
	"fmt"
	"vava6/vatools"
)

type AdminUser struct {
	id            int
	uid           string
	password      string
	checks        int
	logins        int
	lastLoginTime int64
}

func (this *AdminUser) AddLogin() {
	this.logins++
}

func (this *AdminUser) GetID() int {
	return this.id
}

func (this *AdminUser) GetUid() string {
	return this.uid
}

func CheckAdminUser(pwd string) (*AdminUser, error) {
	rss, err := DBSave.Querys("*", "admin_users", fmt.Sprintf("password=\"%s\"", pwd))
	if err != nil {
		return nil, err
	}
	if len(rss) != 1 {
		return nil, fmt.Errorf("NULL")
	}
	rs := rss[0]
	return &AdminUser{
		id:       vatools.SInt(rs["id"]),
		uid:      rs["uid"],
		password: rs["password"],
		checks:   vatools.SInt(rs["checks"]),
		logins:   vatools.SInt(rs["logins"]),
	}, nil
}
