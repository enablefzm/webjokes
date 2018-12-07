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
	nowLoadIdx int // 当前要被查找的键值
	nowPage    int
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
	if len(this.arrJokes) <= idx {
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

	rss, err := this.loadDb()
	if err != nil {
		return
	}
	il := len(rss)
	if il < 1 {
		this.nowLoadIdx = 0
		rss, err = this.loadDb()
		if err != nil {
			return
		}
	}
	il = len(rss)
	this.arrJokes = make([]*Joke, 0, il)
	for i := 0; i < il; i++ {
		rs := rss[i]
		ptJoke := NewJokeOnRs(rs)
		this.arrJokes = append(this.arrJokes, ptJoke)
	}
}

func (this *JokePushPool) loadDb() ([]map[string]string, error) {
	this.nowPage++
	return DBSave.QuerysLimit("*", "joke_text", fmt.Sprintf("is_check > 1 AND is_check < 4 AND push = %d", this.nowLoadIdx), this.nowPage, 100, "id DESC")
}

// 随机获取
func (this *JokePushPool) rndLoadDb() ([]map[string]string, error) {
	// 获取总数量
	rss, err := DBSave.QuerySql("SELECT Count(joke_text.id) AS jokeCounts FROM joke_text WHERE joke_text.is_check > 1 AND joke_text.is_check < 4")
	if err != nil {
		return nil, err
	}
	if len(rss) != 1 {
		return nil, fmt.Errorf("NULL")
	}
	rs := rss[0]
	iCounts := vatools.SInt(rs["jokeCounts"])
	// 获取随机数量
	return nil, fmt.Errorf("TEST")
}
