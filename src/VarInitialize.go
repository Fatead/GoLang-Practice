package main

import "fmt"

/*
 * 变量的初始化，声明变量可以包含初始值，每个变量对应一个
 * 如果初始化值已存在，则可以省略类型，变量会从初始值中获得类型
 */
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "some"
	fmt.Println(i,j)
	fmt.Println(c,python,java)
}
