package models

import (
	"fmt"
	"time"
	"vava6/pools"
	"vava6/vatools"
)

var OBShareJokePool *JokeSharePool

func init() {
	OBShareJokePool = &JokeSharePool{
		sharePool: pools.NewPool(func(id interface{}) (pools.IFObject, error) {
			if jokeID, ok := id.(int); ok {
				return NewShareJokeNoCreate(jokeID)
			}
			return nil, fmt.Errorf("ID Error")
		}),
	}
}

type JokeSharePool struct {
	sharePool *pools.Pool
}

// 查看分享的ID
func (this *JokeSharePool) SeeJoke(jokeID int) (*Joke, error) {
	p, err := this.Get(jokeID)
	if err != nil {
		return nil, err
	}
	p.seeValue++
	return p.Joke, nil
}

// 分享笑话
func (this *JokeSharePool) ShareJoke(jokeID int) error {
	// 判断当前池子里是否有这个ID
	if !OBPushJokePool.CheckIsExist(jokeID) {
		return fmt.Errorf("NO EXIST")
	}
	// 从缓存池里获取对象
	ptJoke, err := this.Get(jokeID)
	if err != nil {
		if err.Error() == "NULL" {
			// 从库里创建
			ptJokeSource, err := NewJokeSourceOnID(jokeID)
			if err != nil {
				return err
			}
			// 生成一个新的JokeShare对象
			ptJoke = &JokeShare{
				JokeSource: ptJokeSource,
			}
			// 新增保存
			errSave := ptJoke.Save()
			if errSave != nil {
				fmt.Println(errSave.Error())
			}
			// 把这个对象放到池子里
			this.sharePool.Put(jokeID, ptJoke)
		} else {
			return err
		}
	}
	if ptJoke != nil {
		ptJoke.shareValue++
	}
	return nil
}

func (this *JokeSharePool) Get(jokeID int) (*JokeShare, error) {
	p, err := this.sharePool.Get(jokeID)
	if err != nil {
		return nil, err
	}
	if ptJoke, ok := p.(*JokeShare); ok {
		return ptJoke, nil
	}
	return nil, fmt.Errorf("TYPE_ERROR")
}

func (this *JokeSharePool) Save() error {
	return this.sharePool.Save()
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
