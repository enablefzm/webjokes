package models

import (
	"fmt"
	"sync"
	"time"
	"vava6/vatools"
)

var OBShareJokePool *JokeSharePool

func init() {
	OBShareJokePool = &JokeSharePool{
		lk:      new(sync.RWMutex),
		mpCache: make(map[int]*JokeShare, 100),
	}
}

type JokeSharePool struct {
	lk      *sync.RWMutex
	mpCache map[int]*JokeShare
}

// 分享笑话
func (this *JokeSharePool) ShareJoke(jokeID int) error {
	var err error
	this.lk.RLock()
	ptShareJoke, ok := this.mpCache[jokeID]
	this.lk.RUnlock()
	if !ok {
		this.lk.Lock()
		ptShareJoke, ok = this.mpCache[jokeID]
		if !ok {
			// 从库里新建一个
			ptShareJoke, err = NewShareJoke(jokeID)
			if err == nil {
				this.mpCache[ptShareJoke.JokeSource.id] = ptShareJoke
			}
		}
		this.lk.Unlock()
	}
	// 新增一次分享记录
	if ptShareJoke != nil {
		ptShareJoke.shareValue++
	}
	return err
}

func NewShareJoke(jokeID int) (*JokeShare, error) {
	// 判断当前库里是否有这个对象
	ptJoke, err := NewShareJokeNoCreate(jokeID)
	if err != nil {
		// 如果返回值是NULL说明数据库里没有这个对象则创建
		if err.Error() == "NULL" {
			ptJokeSource, err := NewJokeSourceOnID(jokeID)
			if err != nil {
				return nil, err
			}
			// 生成一个新的JokeShare对象
			ptJoke = &JokeShare{
				JokeSource: ptJokeSource,
			}
			// 新增保存
			ptJoke.Save()
		}
	}
	return ptJoke, nil
}

// 创建一个不需
func NewShareJokeNoCreate(jokeID int) (*JokeShare, error) {
	rss, err := DBSave.Querys("*", "joke_share", fmt.Sprintf("joke_id=%d", jokeID))
	if err != nil {
		return nil, err
	}
	if len(rss) != 1 {
		return nil, fmt.Errorf("NULL")
	}
	return newShareJokeOnRs(rss[0])
}

func newShareJokeOnRs(rs map[string]string) (*JokeShare, error) {
	jokeID := vatools.SInt(rs["joke_id"])
	if jokeID < 1 {
		return nil, fmt.Errorf("NULL")
	}
	// 获取笑话段子
	ptJoke, err := NewJokeSourceOnID(jokeID)
	if err != nil {
		return nil, err
	}
	return &JokeShare{
		JokeSource: ptJoke,
		id:         vatools.SInt(rs["id"]),
		shareValue: vatools.SInt(rs["shareValue"]),
		seeValue:   vatools.SInt(rs["seeValue"]),
		lastSee:    time.Now().Unix(),
	}, nil
}

type JokeShare struct {
	*JokeSource
	id         int
	shareValue int
	seeValue   int
	lastSee    int64
}

func (this *JokeShare) SetLastSee() {
	this.lastSee = time.Now().Unix()
}

func (this *JokeShare) Save() error {
	if this.id == 0 {
		// 新增
		result, err := DBSave.Insert("joke_share", map[string]interface{}{
			"joke_id":    this.JokeSource.id,
			"shareValue": this.shareValue,
			"seeValue":   this.seeValue,
		})
		if err != nil {
			return err
		}
		autoID, _ := result.LastInsertId()
		this.id = int(autoID)
	} else {
		// 更新
		_, err := DBSave.Update("joke_share", map[string]interface{}{
			"shareValue": this.shareValue,
			"seeValue":   this.seeValue,
		}, map[string]interface{}{"id": this.id})
		if err != nil {
			return err
		}
	}
	return nil
}
