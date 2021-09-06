package main

import "fmt"

//fibonacci 用闭包的机制来实现斐波拉契数列
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
