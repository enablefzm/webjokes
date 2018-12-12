package models

import (
	"fmt"
	"testing"
)

func TestJokes(t *testing.T) {
	i := 10
	v := true
	if v {
		var ii int
		i, ii = createVal()
		fmt.Println(i, ii)
	}
	fmt.Println(i)
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
