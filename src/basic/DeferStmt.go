package main

import "fmt"

func main() {
	// defer语句会将函数推迟到外层函数返回之后执行，推迟调用的函数
	// 其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	defer fmt.Println("world")
	fmt.Println("hello")
	// defer栈
	// 推迟的函数调用会被压入到一个栈中，当外层的函数返回的时候，被推迟的函数会按照
	// 先进后出的顺序调用
	fmt.Println("counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}
