package main

import (
	"fmt"
	"net/http"
)

/*
创建处理器函数，通过处理器获取请求行里面的信息
一个HTTP请求中的请求行主要包装到Request.URL这个结构体里面了
请求体中的内容包装到Request.Header这个结构体中
*/
func handler(w http.ResponseWriter, r *http.Request) {
	//通过Request获取请求行里面的信息
	fmt.Fprintln(w, "你发送的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "查询的地址是：", r.URL)
	fmt.Fprintln(w, "请求的查询字符串为：", r.URL.RawQuery)
	//通过Request获取请求头里面的信息
	fmt.Fprintln(w, "请求头的信息为", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头的信息Accept-Encoding为", r.Header.Get("Accept-Encoding"))
	fmt.Fprintln(w, "请求头的信息Connection为", r.Header.Get("Connection"))
	fmt.Fprintln(w, "请求头的信息Cache-Control为", r.Header.Get("Cache-Control"))

	//获取HTTP请求体中的内容
	//获取请求体内容的长度
	len := r.ContentLength
	//创建byte类型的切片，用于接收请求体中的内容
	body := make([]byte, len)
	//将请求体中的内容读取到切片中
	r.Body.Read(body)
	fmt.Fprintln(w, "请求体中的内容有 ： ", string(body))

	//解析表单，在调用r.Form之前必须执行该操作
	//因为body中的数据是数据流，所以起那面Read(body)后，body中的数据就读不到了，只能读一次
	r.ParseForm()
	fmt.Fprintln(w, "请求参数有", r.Form)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
