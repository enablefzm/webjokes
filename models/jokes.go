package models

import (
	"fmt"
	"vava6/vatools"
)

var prepareJokes []*JokeSource = make([]*JokeSource, 0, 30)

func NewJokeOnRs(rs map[string]string) *Joke {
	return &Joke{
		id:          vatools.SInt(rs["id"]),
		content:     rs["content"],
		keywords:    rs["keywords"],
		vote:        vatools.SInt(rs["vote"]),
		commint:     vatools.SInt(rs["commint"]),
		dateTimeStr: rs["dateTime"],
	}
}

func NewJokeSourceOnRs(rs map[string]string) *JokeSource {
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

type Joke struct {
	id          int
	content     string
	keywords    string
	vote        int
	commint     int
	dateTimeStr string
}

type JokeSource struct {
	*Joke
	isCheck  int8
	checkIds string
}
