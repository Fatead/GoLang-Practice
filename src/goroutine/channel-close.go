package main

import "fmt"

// 关闭管道后不可写入，但是还可以进行读取
func main() {
	initChan := make(chan int, 3)
	initChan <- 100
	initChan <- 200
	close(initChan)
	fmt.Println(<-initChan)
}
