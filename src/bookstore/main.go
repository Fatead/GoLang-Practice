package main

import (
	"net/http"
)

//main 在线商城程序的入口函数
func main() {
	http.ListenAndServe(":8080", nil)
}
