package models

import (
	"fmt"
	"vava6/vatools"
)

type Joke struct {
	id          int
	content     string
	keywords    string
	vote        int
	commint     int
	dataTime    int64
	dataTimeStr string
}

type JokeSource struct {
	*Joke
	isCheck  int8
	checkIds string
}

func NewJokeOnRs(rs map[string]string) *Joke {
	return &Joke{
		id:          vatools.SInt(rs["id"]),
		content:     rs["content"],
		keywords:    rs["keywords"],
		vote:        vatools.SInt(rs["vote"]),
		commint:     vatools.SInt(rs["commint"]),
		dataTime:    vatools.SInt64(rs["dateTime"]),
		dataTimeStr: rs["dateTime"],
	}
}

func NewJokeSourceOnRs(rs map[string]string) *JokeSource {
	fmt.Println(rs)
	return &JokeSource{
		Joke:     NewJokeOnRs(rs),
		isCheck:  int8(vatools.SInt(rs["is_check"])),
		checkIds: rs["check_ids"],
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
	ptJokesSource := NewJokeSourceOnRs(rs)
	return ptJokesSource, nil
}
