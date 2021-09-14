package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:15927029790@tcp(localhost:3306)/sakila")
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
	defer db.Close()

	sqlStr := "insert into actor(actor_id, first_name, last_name) values (?, ?, ?)"
	//预编译
	inStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
	}
	//执行语句
	_, err2 := inStmt.Exec(1245, "tomcat", "tomcat")
	if err2 != nil {
		fmt.Println("执行时出现异常")
	}
}
