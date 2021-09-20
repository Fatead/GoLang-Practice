package test

import (
	"awesomeProject/src/bookstore/dao"
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
	"testing"
)

func TestAddSession(t *testing.T) {
	s := &model.Session{
		SessionId: utils.CreateUUID(),
		Username:  "Jack",
		UserId:    16,
	}
	err := dao.AddSession(s)
	if err != nil {
		fmt.Println("添加Session过程中出现异常", err)
	}
}

func TestDeleteSession(t *testing.T) {
	session_id := ""
	affect, err := dao.DeleteSessionById(session_id)
	if err != nil {
		fmt.Println("删除session过程中出现异常:", err)
	} else {
		fmt.Println("受影响的行数为：", affect)
	}
}
