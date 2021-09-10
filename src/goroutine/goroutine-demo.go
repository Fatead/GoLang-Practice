package main

import (
	"fmt"
	"time"
)

/*
 * Go语言的携程Goroutine
 * go f(x, y, z)会启动一个新的协程并执行
 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("wrold")
	say("hello")
}
