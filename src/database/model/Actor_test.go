package model

import (
	"fmt"
	"testing"
)

func TestAddActor(t *testing.T) {
	fmt.Println("测试添加用户...")
	actor := &Actor{}
	if actor != nil {
		err := actor.AddActor()
		if err != nil {
			t.Log("测试添加失败")
			return
		}
		t.Log("测试添加成功")
	}
}
