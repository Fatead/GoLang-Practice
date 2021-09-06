package main

import "fmt"

//常量可以是字符、字符串、布尔值或数值，常量不能用:=语法声明

const Pi = 3.14

func main() {
	fmt.Println(Pi)
	const world = "世界"
	fmt.Println("Hello", world)
}
