package models

import (
	"fmt"
	"testing"
)

func TestJokes(t *testing.T) {
	rss, _ := OBPushJokePool.rndLoadDb2()
	// t.Log(rss)
	for _, rs := range rss {
		fmt.Println(rs["id"])
	}
	// t.Log(len(rss))
}
