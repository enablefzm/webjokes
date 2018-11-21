package models

import (
	"fmt"
	"testing"
)

func TestJokes(t *testing.T) {
	//	rss, err := DBSave.QuerysLimit("*", "joke_text", "is_check=0", 1, 2, "id DESC")
	//	if err != nil {
	//		t.Log(err.Error())
	//	} else {
	//		for _, rs := range rss {
	//			t.Log(rs)
	//		}
	//	}
	fmt.Println("count:", OBJokePool.Count())
	for i := 0; i < 61; i++ {
		p, _ := OBJokePool.Get()
		// fmt.Println(p)
		fmt.Println(p.id, p.content)
	}
	fmt.Println("count:", OBJokePool.Count())
}
