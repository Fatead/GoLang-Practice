package model

import (
	"awesomeProject/src/database/utils"
	"fmt"
)

type Actor struct {
	ActorId    int
	FirstName  string
	LastName   string
	LastUpdate string
}

func (actor *Actor) AddActor() error {
	//sql语句
	sqlStr := "insert into actor(actor_id, first_name, last_name) values (?, ?, ?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
	}
	//执行语句
	_, err2 := inStmt.Exec(336, "tomcat", "tomcat")
	if err2 != nil {
		fmt.Println("执行时出现异常")
		return err2
	}
	return nil
}

// 根据用户的id从数据库中查询一条记录
func (actor *Actor) GetUserById() (*Actor, error) {
	sqlStr := "select first_name,last_name from actor where actor_id = ?"
	row := utils.Db.QueryRow(sqlStr, 12)
	var firstName string
	var secondName string
	err := row.Scan(&firstName, &secondName)
	if err != nil {
		fmt.Println("查询数据出错")
	}
	return nil, err
}
