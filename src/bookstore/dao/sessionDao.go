package dao

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
	"net/http"
)

//AddSession 向数据库中添加一个session信息
func AddSession(s *model.Session) error {
	sqlStr := "insert into sessions (session_id, username, user_id) values (?, ?, ?)"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译语句出错")
		return err
	}
	//使用MD5对session_id进行加密，为了防止攻击者伪造session 信息，session id储存在数据库之前也需要先加密
	s.SessionId = utils.Md5(s.SessionId)
	_, err = stmt.Exec(&s.SessionId, &s.Username, &s.UserId)
	if err != nil {
		fmt.Println("插入语句执行出错")
		return err
	}
	return nil
}

//DeleteSessionById 根据ID从数据库中删除对应的Session
func DeleteSessionById(session_id string) (int64, error) {
	sqlStr := "delete from sessions where session_id = ?"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译语句出错")
		return 0, err
	}
	session_id = utils.Md5(session_id)
	res, err := stmt.Exec(session_id)
	if err != nil {
		fmt.Println("执行出错,err:", err)
		return 0, err
	}
	affect, errRes := res.RowsAffected()
	if errRes != nil {
		fmt.Println("取出受影响行数时出现异常，err:", errRes)
		return 0, errRes
	}
	return affect, nil
}

//IsLogin 根据HTTP请求报文信息来验证用户是否登录过，如果已经登录过，那么直接根据HTTP请求报文中的Cookie中的Session_id获取该用户对应的Session并返回
func IsLogin(r *http.Request) (bool, *model.Session, error) {
	// 根据Cookie信息检查用户是否已经登录
	cookie, _ := r.Cookie("user")
	if cookie == nil {
		fmt.Println("获取Cookie失败，Cookie可能不存在，还未登录")
		return false, nil, nil
	} else {
		sqlStr := "select session_id, user_id, username from sessions where session_id = ?"
		session_id := utils.Md5(cookie.Value)
		session := &model.Session{}
		err := utils.Db.QueryRow(sqlStr, session_id).Scan(&session.SessionId, &session.UserId, &session.Username)
		if err != nil {
			fmt.Println("数据库中没有查找到该Session相关的记录信息，err: ", err)
			return false, nil, nil
		}
		return true, session, nil
	}
}
