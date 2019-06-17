package models

import (
	"vava6/vaini"
	"vava6/vatools"
)

// 获取需要审核的对象
var OBJokePool *JokePool = NewJokePool(func() ([]map[string]string, error) {
	// 获取条件判断信息
	strSql := "is_check=0 AND vote >= 2000"
	page := 60
	desc := "id DESC"
	// 读取ini文件
	c := vaini.NewConfig("cfg.ini")
	if mpCfg, ok := c.GetNode("JOKE_CHECK"); ok {
		for k, v := range mpCfg {
			switch k {
			case "where":
				strSql = v
			case "page":
				page = vatools.SInt(v)
			case "desc":
				desc = v
			}
		}
	}
	return DBSave.QuerysLimit("*", "joke_text", strSql, 1, page, desc)
})

type fnGetJokes func() ([]map[string]string, error)

type JokePool struct {
	arrJoke   []*JokeSource
	fGetJokes fnGetJokes
}

func NewJokePool(fn fnGetJokes) *JokePool {
	p := &JokePool{
		arrJoke:   make([]*JokeSource, 0, 30),
		fGetJokes: fn,
	}
	return p
}

func (this *JokePool) loadJokes() int {
	rss, err := this.fGetJokes()
	if err != nil {
		return 0
	}
	il := len(rss)
	if il < 1 {
		return il
	}
	this.arrJoke = make([]*JokeSource, 0, 60)
	for _, rs := range rss {
		this.arrJoke = append(this.arrJoke, NewJokeSourceOnRs(rs))
	}
	return il
}

func (this *JokePool) Get() (*JokeSource, bool) {
	ptJoke, ok := this.getJoke()
	if !ok {
		// 重新装载对象
		resValue := this.loadJokes()
		if resValue < 1 {
			return nil, false
		}
		ptJoke, ok = this.getJoke()
	}
	return ptJoke, ok
}

func (this *JokePool) getJoke() (*JokeSource, bool) {
	var k int
	var ptJoke *JokeSource
	for k, ptJoke = range this.arrJoke {
		if ptJoke != nil {
			this.arrJoke[k] = nil
			return ptJoke, true
		}
	}
	return nil, false
}

func (this *JokePool) Count() int {
	i := 0
	for _, p := range this.arrJoke {
		if p != nil {
			i++
		}
	}
	return i
}

func init() {
	OBJokePool.loadJokes()
}
