package models

import (
	"fmt"
	"vava6/vatools"
)

const (
	TABLE_ADMIN_USER = "admin_users"
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
	if err := this.reload(); err != nil {
		return
	}
	this.logins++
	DBSave.Update(TABLE_ADMIN_USER, map[string]interface{}{
		"logins":          this.logins,
		"last_login_time": vatools.GetNowTimeString(),
	}, map[string]interface{}{"id": this.id})
}

func (this *AdminUser) AddCheck() {
	if err := this.reload(); err != nil {
		return
	}
	this.checks++
	DBSave.Update(TABLE_ADMIN_USER, map[string]interface{}{"checks": this.checks}, map[string]interface{}{"id": this.id})
}

func (this *AdminUser) GetID() int {
	return this.id
}

func (this *AdminUser) GetUid() string {
	return this.uid
}

func (this *AdminUser) reload() error {
	rss, err := DBSave.Querys("checks, logins", TABLE_ADMIN_USER, fmt.Sprintf("id=%d", this.id))
	if err != nil {
		return err
	}
	if len(rss) != 1 {
		return fmt.Errorf("NULL")
	}
	rs := rss[0]
	this.checks = vatools.SInt(rs["checks"])
	this.logins = vatools.SInt(rs["logins"])
	return nil
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
