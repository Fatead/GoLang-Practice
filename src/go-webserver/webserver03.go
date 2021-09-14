package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler2 struct {
}

func (m *MyHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过详细配置服务器的信息来处理请求")
}

func main() {
	myHandler2 := MyHandler2{}
	//通过http.Server来配置详细的服务器信息
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler2,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}
