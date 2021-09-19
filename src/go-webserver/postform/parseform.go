package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	/*
	 * 使用PostForm只获取请求体中的属性，而不获取请求行中的属性
	 * 但是PostForm只支持application/x-www-form-urlencoded编码
	 * 如果form表单中的enctype属性值为multipart/form-data，那么需要使用MultipartForm字段来获取表单中的数据
	 */
	fmt.Fprintln(w, "POST请求中的form表单中的请求参数有：", r.PostForm)
}

func main() {
	http.HandleFunc("/form", handler)
	http.ListenAndServe(":8080", nil)
}
