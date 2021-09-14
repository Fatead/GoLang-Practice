package main

import (
	"fmt"
	"net/http"
)

//实现Go Web -> Hello World
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Hello World", r.URL.Path)
}

func main() {
	//handlerFunc会将handler转换成一个实现了Handler接口的HandlerFunc类型
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
