package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

//初始化数据库连接
func init() {
	Db, err = sql.Open("mysql", "root:15927029790@tcp(localhost:3306)/bookstore?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
}
