package models

import (
	"fmt"
	"vava6/vatools"
)

const (
	CHECK_NO   int8 = 1
	CHECK_OK   int8 = 2
	CHECK_GOOD int8 = 3
)

// var prepareJokes []*JokeSource = make([]*JokeSource, 0, 30)

func NewJokeOnRs(rs map[string]string) *Joke {
	return &Joke{
		id:          vatools.SInt(rs["id"]),
		content:     rs["content"],
		keywords:    rs["keywords"],
		vote:        vatools.SInt(rs["vote"]),
		comment:     vatools.SInt(rs["comment"]),
		dateTimeStr: rs["dateTime"],
	}
}

func NewJokeSourceOnRs(rs map[string]string) *JokeSource {
	return &JokeSource{
		Joke:        NewJokeOnRs(rs),
		isCheck:     int8(vatools.SInt(rs["is_check"])),
		checkIds:    rs["check_ids"],
		sourceTable: rs["sourceTable"],
		labels:      rs["labels"],
	}
}

func NewJokeSourceOnID(id int) (*JokeSource, error) {
	rss, err := DBSave.Querys("*", "joke_text", fmt.Sprintf("id=%d", id))
	if err != nil {
		return nil, err
	}
	if len(rss) != 1 {
		return nil, fmt.Errorf("NULL")
	}
	rs := rss[0]
	rs["sourceTable"] = "joke_text"
	ptJokesSource := NewJokeSourceOnRs(rs)
	return ptJokesSource, nil
}

type Joke struct {
	id          int
	content     string
	keywords    string
	vote        int
	comment     int
	dateTimeStr string
}

func (this *Joke) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"id":          this.id,
		"content":     this.content,
		"vote":        this.vote,
		"comment":     this.comment,
		"dateTimeStr": this.dateTimeStr,
	}
}

type JokeSource struct {
	*Joke
	isCheck     int8
	checkIds    string
	sourceTable string
	labels      string
}

func (this *JokeSource) SetCheckState(v int8) {
	this.isCheck = v
}

func (this *JokeSource) SetCheckIDs(id int) {
	if len(this.checkIds) < 1 {
		this.checkIds = fmt.Sprintf("%d", id)
	} else {
		this.checkIds = fmt.Sprintf("%s,%d", this.checkIds, id)
	}
}

func (this *JokeSource) SetLabels(strLabels string) {
	this.labels = strLabels
}

func (this *JokeSource) Updata() {
	DBSave.Update(this.sourceTable, map[string]interface{}{
		"is_check":  this.isCheck,
		"check_ids": this.checkIds,
		"labels":    this.labels,
	}, map[string]interface{}{"id": this.id})
}

func (this *JokeSource) GetInfo() map[string]interface{} {
	mpInfo := this.Joke.GetInfo()
	mpInfo["isCheck"] = this.isCheck
	mpInfo["checkIds"] = this.checkIds
	return mpInfo
}
