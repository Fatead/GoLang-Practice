package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t, _ := template.ParseFiles("C:\\Users\\Zero\\GolandProjects\\awesomeProject\\src\\go-webserver\\template\\Hello.html")
	//执行模板
	t.Execute(w, "Hello World")
}

func main() {
	http.HandleFunc("/template", handler)
	http.ListenAndServe(":8080", nil)
}
