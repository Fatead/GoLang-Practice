package main

import (
	"awesomeProject/src/go-kit/v1/endpoint_v1"
	"awesomeProject/src/go-kit/v1/service_v1"
	"awesomeProject/src/go-kit/v1/transport_v1"
	"fmt"
	"net/http"
)

func main() {
	server := service_v1.NewService()
	endpoints := endpoint_v1.NewEndPointServer(server)
	httpHandler := transport_v1.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
