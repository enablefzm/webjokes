package controllers

import (
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
	act := this.GetString("act")
	switch act {
	// 收到分享ID
	case "share":
		id := vatools.SInt(this.GetString("id"))
		err := models.OBShareJokePool.ShareJoke(id)
		msg := ""
		result := 0
		if err != nil {
			result = -1
			msg = err.Error()
		}
		this.OutJson(map[string]interface{}{
			"reuslt": result,
			"info":   msg,
		})
	// 查看分享
	case "seeshare":
		id := vatools.SInt(this.GetString("id"))
		ptJoke, err := models.OBShareJokePool.SeeJoke(id)
		if err != nil {
			this.OutJson(map[string]interface{}{
				"result": -1,
				"info":   err.Error(),
			})
		} else {
			this.OutJson(map[string]interface{}{
				"result": 0,
				"info":   ptJoke.GetInfo(),
			})
		}
	default:
		idx := vatools.SInt(this.GetString("jid"))
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
}

func (this *GetJokeControllers) OutJson(info interface{}) {
	this.Data["json"] = info
	this.ServeJSON()
}
