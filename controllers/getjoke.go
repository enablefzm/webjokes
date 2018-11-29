package controllers

import (
	"fmt"
	"vava6/vatools"
	"webjokes/models"

	"github.com/astaxie/beego"
)

const (
	WX_USER = "WX_USER"
	WX_IDX  = "WX_IDX"
)

type ResultDb struct {
	Result int8        `json:"result"`
	Info   interface{} `json:"info"`
}

type GetJokeControllers struct {
	beego.Controller
}

func (this *GetJokeControllers) Get() {
	this.doing()
}

func (this *GetJokeControllers) Post() {
	this.doing()
}

func (this *GetJokeControllers) doing() {
	//	idx := 0
	//	// 获取当前索引
	//	v := this.GetSession(WX_IDX)
	//	if v != nil {
	//		var ok bool
	//		idx, ok = v.(int)
	//		if !ok {
	//			idx = 0
	//		}
	//	}
	idx := vatools.SInt(this.GetString("jid"))
	fmt.Println(this.GetString("jid"), idx)
	res, err := models.OBPushJokePool.GetJoke(idx)
	if err != nil {
		this.OutJson(map[string]interface{}{
			"result": -1,
			"info":   err.Error(),
		})
	} else {
		this.OutJson(map[string]interface{}{
			"result": 0,
			"info":   res.PtJoke.GetInfo(),
			"nextId": res.NextIdx,
		})
	}
}

func (this *GetJokeControllers) OutJson(info interface{}) {
	this.Data["json"] = info
	this.ServeJSON()
}
