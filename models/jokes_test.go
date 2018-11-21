package models

import (
	"testing"
)

func TestJokes(t *testing.T) {
	pt, err := NewJokeSourceOnID(97743014)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else {
		t.Log(pt.dateTimeStr)
	}

	arr := []string{"1"}
	t.Log(arr[1:])
}
