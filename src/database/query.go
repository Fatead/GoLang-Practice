package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 执行数据库的查询请求
func main() {
	db, err := sql.Open("mysql", "root:15927029790@tcp(localhost:3306)/sakila")
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
	defer db.Close()
	sqlStr := "select actor_id, first_name, last_name from actor"
	rows, err := db.Query(sqlStr)
	for rows.Next() {
		var actorId int
		var firstName string
		var secondName string
		err := rows.Scan(&actorId, &firstName, &secondName)
		if err != nil {
			fmt.Println("query error")
		}
		fmt.Println("id :", actorId, "  firstName:", firstName, "  secondName", secondName)
	}
}
