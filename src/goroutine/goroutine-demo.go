package main

import (
	"fmt"
	"time"
)

/*
 * Go语言的携程Goroutine
 * go f(x, y, z)会启动一个新的协程并执行
 * 协程也称用户态的线程或者轻量级的线程，由用户进行调度
 * 有自己独立的栈空间，共享堆空间
 * 如果主线程退出了，即使协程还没有执行完毕，也会退出
 * 协程也可以在主线程没有退出前，就自己结束了
 *
 * 主线程是一个物理线程，直接作用在cpu上的，是重量级的，非常耗费CPU资源，
 * 协程是从主线程开启的，轻量级的线程，是逻辑态，对资源的消耗相对较小
 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("wrold")
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
		time.Sleep(100 * time.Millisecond)
	}
}
