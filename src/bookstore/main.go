package main

import (
	"awesomeProject/src/bookstore/controller"
	"net/http"
)

//main 在线商城程序的入口函数
func main() {
	//配置静态资源的映射规则
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("src/bookstore/views/static/"))))

	http.Handle("/pages/",
		http.StripPrefix("/pages/", http.FileServer(http.Dir("src/bookstore/views/pages/"))))

	//配置用户相关的处理请求
	http.HandleFunc("/main", controller.IndexHandler)
	http.HandleFunc("/toLogin", controller.ToLogin)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/FindUserByName", controller.FindUserByName)

	//配置图书相关的处理请求

	http.ListenAndServe(":8080", nil)
}
