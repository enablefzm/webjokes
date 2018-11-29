package models

import (
	"vava6/vatools"
	"fmt"
)

type ResJokePush struct {
	PtJoke *Joke
	NextIdx int
}

type JokePushPool struct {
	isRuning   bool
	nowLoadIdx int		// 当前要被查找的键值
	arrJokes   []*Joke
}

func (this *JokePushPool) GetJoke(idx int) (ResJokePush, error) {
	if  idx >= len(this.arrJokes) {
		// 判断是否需要执行更新
		idx = 0
	}
	return ResJokePush{
		PtJoke: this.arrJokes[idx],
		NextIdx: idx++,
	}
}

func (this *JokePushPool) Run() {
	if this.isRuning {
		return
	}
}

func (this *JokePushPool) _running() {
	this.load();
}

func (this *JokePushPool) load() {
	// 获取随机总理
	rss, err := DBSave.QuerySql(fmt.Sprintf("SELECT Count(joke_text.id) AS counts FROM joke_text WHERE is_check > 1 AND is_check < 4 AND push = %d", this.nowLoadIdx))
	if err != nil || len(rss) < 1{
		return
	}
	rs := rss[0]
	count := vatools.SInt(rs["counts"])
	rss, err := DBSave.QuerysLimit("*", "joke_text", fmt.Sprintf("is_check > 1 AND is_check < 4 AND push = %d", this.nowLoadIdx), 1, 100, "id DESC")
}
