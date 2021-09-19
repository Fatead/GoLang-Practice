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

/*
 *  init()函数会在每个包完成初始化后自动执行，并且执行优先级比main函数高
 *  init函数通常被用来：
 *  1. 对变量进行初始化
 *  2. 检查、修复程序的状态
 *  3. 注解
 *  4. 运行一次计算
 */

func init() {
	db, err := sql.Open("mysql", "root:15927029790@tcp(localhost:3306)/sakila")
	if err != nil {
		fmt.Println("连接数据库失败...")
		panic(err.Error())
	}
	fmt.Println("数据库连接成功")
	Db = db
}
