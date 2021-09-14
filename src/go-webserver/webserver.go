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
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	//myHandler := MyHandler{}
	//http.Handle("/myHandler", &myHandler)
	http.ListenAndServe(":8080", mux)
}
