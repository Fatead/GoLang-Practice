package main

import "fmt"

//结构体文法，通过直接列出字段的值来新分配一个结构体

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} //创建一个Vertex类型的结构体
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2} // 创建一个指向*Vertex类型的指针
)

func main() {
	v4 := Vertex{}
	p.X = 4
	fmt.Println(v1, v2, v3, *p, v4)
}
