package models

import (
	"testing"
	"time"
)

func TestJokes(t *testing.T) {
	pt, err := NewJokeSourceOnID(97743014)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else {
		t.Log(pt.dataTime)
		t.Log(pt.dataTimeStr)
		// t.Log(time.ParseDuration(pt.dataTimeStr))
		t.Log(time.Parse("", pt.dataTimeStr))
	}
}
