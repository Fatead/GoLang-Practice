package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败")
	}
	fmt.Println("连接成功")
	conn.LocalAddr()
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入【终端】
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string error", err)
		}
		n, err := conn.Write([]byte(line)) //将string强制转换为byte类型的切片
		if err != nil {
			fmt.Println("write error")
		}
		fmt.Println(n)
	}
}
