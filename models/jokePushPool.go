package models

import (
	"fmt"
	"sync"
	"time"
	"vava6/vatools"
)

const (
	REFRESH_TIME int64 = 600
)

var OBPushJokePool *JokePushPool = NewJokePushPool()

type ResJokePush struct {
	PtJoke  *Joke
	NextIdx int
}

func NewJokePushPool() *JokePushPool {
	return &JokePushPool{
		arrJokes: make([]*Joke, 0, 1),
		lk:       new(sync.RWMutex),
	}
}

type JokePushPool struct {
	isRuning   bool
	nowLoadIdx int     // 当前要被查找的键值
	arrJokes   []*Joke // 数组里存放要被推送给段友的段子
	nowTime    int64   // 计录当前刷新的时间
	lk         *sync.RWMutex
}

func (this *JokePushPool) GetJoke(idx int) (ResJokePush, error) {
	il := len(this.arrJokes)
	if idx >= il {
		// 判断是否需要执行更新
		nowTimestamp := time.Now().Unix()
		if (nowTimestamp - this.nowTime) > REFRESH_TIME {
			// 执行刷新
			this.lk.Lock()
			if (nowTimestamp - this.nowTime) > REFRESH_TIME {
				// 刷新数据
				this.load()
			}
			this.lk.Unlock()
			idx = 0
		} else {
			// 随机获取一个笑话
			idx = vatools.CRnd(0, il-1)
		}

	}
	if len(this.arrJokes) < idx {
		return ResJokePush{PtJoke: nil, NextIdx: idx}, fmt.Errorf("NULL")
	}
	nextIdx := idx + 1
	return ResJokePush{
		PtJoke:  this.arrJokes[idx],
		NextIdx: nextIdx,
	}, nil
}

func (this *JokePushPool) load() {
	this.nowTime = time.Now().Unix()
	//	strSql := fmt.Sprintf("SELECT Count(joke_text.id) AS counts FROM joke_text WHERE is_check > 1 AND is_check < 4 AND push = %d", this.nowLoadIdx)
	//	// 获取随机总理
	//	rss, err := DBSave.QuerySql(strSql)
	//	if err != nil || len(rss) < 1 {
	//		return
	//	}
	//	rs := rss[0]
	//	count := vatools.SInt(rs["counts"])
	rss, err := DBSave.QuerysLimit("*", "joke_text", fmt.Sprintf("is_check > 1 AND is_check < 4 AND push = %d", this.nowLoadIdx), 1, 100, "id DESC")
	if err != nil {
		return
	}
	il := len(rss)
	this.arrJokes = make([]*Joke, 0, il)
	for i := 0; i < il; i++ {
		rs := rss[i]
		ptJoke := NewJokeOnRs(rs)
		this.arrJokes = append(this.arrJokes, ptJoke)
	}
	fmt.Println("【加载新段子】")
}
