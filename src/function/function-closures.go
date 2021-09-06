package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 函数的闭包，Go函数可以是一个闭包，闭包是一个函数值，它引用了其函数体之外的变量，
// 该函数可以访问并赋予其引用变量的值，换句话说，该函数被这些变量绑定在一起
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

}
