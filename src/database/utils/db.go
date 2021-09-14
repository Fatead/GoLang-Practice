package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err := sql.Open("mysql", "root:15927029790@tcp(localhost:3306)/sakila")
	if err != nil {
		fmt.Println("连接数据库失败...")
		panic(err.Error())
	}
	Db.Ping()
}
