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
	sqlStr := "insert into actor(actor_id, first_name, last_name, last_update) values (?, ?, ?, ?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
	}
	//执行语句
	_, err2 := inStmt.Exec("12346", "tomcat", "tomcat", "123456")
	if err2 != nil {
		fmt.Println("执行时出现异常")
		return err2
	}
	return nil
}
