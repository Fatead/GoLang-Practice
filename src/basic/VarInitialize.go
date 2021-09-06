package main

import "fmt"

/*
 * 变量的初始化，声明变量可以包含初始值，每个变量对应一个
 * 如果初始化值已存在，则可以省略类型，变量会从初始值中获得类型
 */
var i, j int = 1, 2

func main() {
	// 等价于 c,python,java := true, false,"some"
	var c, python, java = true, false, "some"
	fmt.Println(i, j)
	fmt.Println(c, python, java)
	// 短变量声明，在函数中，简洁赋值语句 := 可以在类型明确的地方代替var声明
	K := 3
	fmt.Println(K)
	i := 42
	fmt.Println(i)
	//强制类型转换
	f := float64(i)
	fmt.Println(f)
	u := uint(f)
	fmt.Println(u)
}
