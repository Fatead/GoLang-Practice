package test

import (
	"awesomeProject/src/bookstore/dao"
	"awesomeProject/src/bookstore/model"
	"fmt"
	"testing"
)

var userinfo = &model.UserInfo{
	ID:       19,
	Username: "Tom",
	Password: "123456",
	Email:    "123456@qq.com",
}

func TestLogin(t *testing.T) {
	flag, _, err := dao.Login("lili", "123456")
	if err != nil {
		fmt.Println("登录出现异常，err:", err)
	}
	if flag {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}

func TestFindByName(t *testing.T) {
	user, err := dao.FindUserByName("lili")
	if err != nil {
		fmt.Println("从数据库查找该用户过程中出现异常")
	} else {
		fmt.Println("查询到了用户的信息,", user)
	}
}

func TestAddUser(t *testing.T) {
	name, err := dao.FindUserByName(userinfo.Username)
	if err == nil && name != nil {
		fmt.Println("该用户信息已经存在，不能重复添加")
		return
	}
	err = dao.AddUser(userinfo)
	if err != nil {
		fmt.Println("添加用户信息过程中出现异常")
	} else {
		fmt.Println("查询用户，验证该用户信息是否已经成功添加")
		byName, err := dao.FindUserByName(userinfo.Username)
		if err != nil {
			fmt.Println("查询用户信息过程中出现了异常")
		} else {
			fmt.Println("成功查询到已经添加的用户的信息：", byName)
		}
	}
}

func TestDeleteUserById(t *testing.T) {
	_, err := dao.DeleteUserById(userinfo.ID)
	if err != nil {
		fmt.Println("删除过程中出现异常")
	} else {
		fmt.Println("查询该用户的信息，验证是否已经成功删除")
		name, err := dao.FindUserById(userinfo.ID)
		if err != nil || name == nil {
			fmt.Println("删除用户成功")
		} else {
			fmt.Println("查询到该用户信息，删除用户失败")
		}
	}
}

func TestUpdatePwdById(t *testing.T) {
	password := "2333"
	_, err := dao.UpdatePwdById(userinfo.ID, password)
	if err != nil {
		fmt.Println("更新用户信息失败")
	} else {
		//用新的用户信息进行登录，看看用户信息是否已经更新成功
		login, uid, err := dao.Login(userinfo.Username, password)
		if err != nil || login == false {
			fmt.Println("使用新的用户信息登录失败，密码没有被正确更新")
		} else {
			fmt.Println("使用新的密码登录成功，用户id为", uid)
		}
	}
}
