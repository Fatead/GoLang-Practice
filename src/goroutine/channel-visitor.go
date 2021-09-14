package main

import "fmt"

func main() {
	channel := make(chan int, 100)
	for i := 0; i < 100; i++ {
		channel <- i
	}
	//遍历管道时需要先关闭管道，否则可能会引起死锁
	close(channel)
	for i := range channel {
		fmt.Println(i)
	}
}
