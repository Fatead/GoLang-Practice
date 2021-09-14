package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送过来的数据
	for {
		byteBuf := make([]byte, 1024)
		//等待客户端通过conn发送消息，如果客户端没有使用write发送数据，那么协程就阻塞在这里
		fmt.Println("服务器在等待客户端的输入")
		n, err := conn.Read(byteBuf)
		if err != nil {
			fmt.Println("读取过程中出现了错误")
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(byteBuf[:n]))
	}
	defer conn.Close() //关闭conn
}

func main() {
	fmt.Println("服务器端开始监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("error ", err)
	}
	defer listen.Close()
	for {
		//等待客户端的连接
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("accept error")
		}
		go process(conn)
	}
}
