package models

import (
	"fmt"
	"testing"
)

func TestJokes(t *testing.T) {
	mp := make(map[int]TmpUser)
	mp[1] = TmpUser{
		name: "jimmy",
		age:  38,
	}

	for k, v := range mp {
		fmt.Println(k, v)
		// v.SetName("jimmyFan")
		p := &v
		fmt.Printf("%p \n", p)
		p.name = "jimmyFan"
		p.SetName("jimmyfan")

		if t, ok := mp[k]; ok {
			t.name = "jimmyfan1"
			fmt.Println("set new")
		}

		fmt.Println(*p)

		//		p := &v
		//		p.name = "jimmyFan"
	}

	fmt.Println(mp)

	if t, ok := mp[1]; ok {
		t.name = "jimmyfan1"
		mp[1] = t
		fmt.Println("set new", mp)
	}
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
