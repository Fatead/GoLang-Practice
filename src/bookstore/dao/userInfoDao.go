package dao

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
)

//FindUserByName 根据用户名查找用户的信息
func FindUserByName(username string) (info *model.UserInfo, err error) {
	sqlStr := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	info = &model.UserInfo{}
	errNotFound := row.Scan(&info.ID, &info.Username, &info.Password, &info.Email)
	if errNotFound != nil {
		return nil, errNotFound
	}
	return info, nil
}

//FindUserById 根据用户id查找用户
func FindUserById(id int) (info *model.UserInfo, err error) {
	sqlStr := "select id,username,password,email from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, id)
	info = &model.UserInfo{}
	errNotFound := row.Scan(&info.ID, &info.Username, &info.Password, &info.Email)
	if errNotFound != nil {
		return nil, errNotFound
	}
	return info, nil
}

//Login 登录，验证密码
func Login(username string, password string) (b bool, uid int, err error) {
	sqlStr := "select id,password from users where username = ?"
	pwd := " "
	b = false
	//从数据库中根据username查出一条数据
	err = utils.Db.QueryRow(sqlStr, username).Scan(&uid, &pwd)
	if err != nil {
		return b, 0, err
	} else {
		//将密码用MD5进行加密后比较是否一致
		code := utils.Md5(password)
		if code == pwd {
			b = true
		} else {
			return b, 0, nil
		}
	}
	return b, uid, nil
}

//AddUser 新增一个用户
func AddUser(info *model.UserInfo) error {
	sqlStr := "insert into users(username, password, email) values (?, ?, ?)"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译 Prepare 出错，err", err)
		return err
	}
	//使用MD5算法对密码加密后入库
	info.Password = utils.Md5(info.Password)
	_, errExec := stmt.Exec(info.Username, info.Password, info.Email)
	if errExec != nil {
		fmt.Println("执行插入语句出错", errExec)
		return errExec
	}
	return nil
}

func DeleteUserById(userId int) (int64, error) {
	sqlStr := "delete from users where id = ?"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译语句出错", err)
		return 0, err
	}
	exec, err := stmt.Exec(userId)
	if err != nil {
		fmt.Println("执行语句出错")
		return 0, err
	}
	affect, errRes := exec.RowsAffected()
	if errRes != nil {
		fmt.Println("取出受影响行数时出现异常，err:", errRes)
		return 0, errRes
	}
	return affect, nil
}

//UpdatePwdById 根据用户的ID修改用户的密码
func UpdatePwdById(id int, password string) (int64, error) {
	sqlStr := "update users set password = ? where id = ?"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译语句出错")
		return 0, err
	}
	password = utils.Md5(password)
	res, err := stmt.Exec(password, id)
	if err != nil {
		fmt.Println("执行密码更新语句出错")
		return 0, err
	}
	affect, errRes := res.RowsAffected()
	if errRes != nil {
		fmt.Println("取出受影响行数时出现异常，err:", errRes)
		return 0, errRes
	}
	return affect, nil
}
