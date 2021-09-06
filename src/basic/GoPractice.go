package main

import "fmt"

// 变量的命名，跟函数的参数列表一样，类型在最后
var c, python, java bool

//add 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其他的都可以忽略
//例如 x int, y int 可以被缩写为x
func add(x, y int) int {
	return x + y
}

//swap 多值返回， 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

//split 命名返回值，Go的返回值可以被命名，它们会被视作定义在函数顶部的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(50, 66))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
}
