package models

import (
	"fmt"
	"testing"
)

func TestJokes(t *testing.T) {

	jokeID := 121289760 // joke.PtJoke.id

	//	fmt.Println("分享ID：", jokeID)
	//	err := OBShareJokePool.ShareJoke(jokeID)
	//	fmt.Println(err)
	//	if err != nil {
	//		fmt.Println("分享发生错误：", err.Error())
	//	}

	// 查看分享
	for i := 0; i < 10; i++ {
		OBShareJokePool.SeeJoke(jokeID)
		fmt.Println(i)
	}
	ptJoke, err := OBShareJokePool.SeeJoke(jokeID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ptJoke.GetInfo())
	}

	errSave := OBShareJokePool.Save()
	t.Log("OK", errSave)

}

func createVal() (int, int) {
	return 9, 8
}

type TmpUser struct {
	name string
	age  int
}

func (this *TmpUser) GetName() string {
	return this.name
}

func (this *TmpUser) SetName(newName string) {
	this.name = newName
}
