package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int
	UserName string
	PassWord string
}

func handler(w http.ResponseWriter, r *http.Request) {
	//设置响应头中内容的类型
	w.Header().Set("Content-Type", "application/json")
	user := User{
		ID:       1,
		UserName: "admin",
		PassWord: "123456",
	}
	//将user数据进行序列化
	json_data, _ := json.Marshal(user)
	w.Write(json_data)
}

func main() {
	http.HandleFunc("/json", handler)
	http.ListenAndServe(":8080", nil)
}
