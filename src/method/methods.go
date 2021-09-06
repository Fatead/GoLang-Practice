package main

import (
	"fmt"
	"math"
)

/*
 * Go没有类，不过可以为结构体类型定义方法
 * Go中方法和函数的定义有所不同
 * 函数是指不属于任何结构体、类型的方法
 * 函数是没有接收者的，而方法是有接收者的
 * 方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者，
 * 这样我们定义的这个方法就和接收者绑定在了一起，称为这个接收者的方法
 */

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	//在绑定之后，可以直接通过接收者来调用相应的方法
	fmt.Println(v.Abs())
}
