package main

/*
可以为非结构体类型声明方法,但是只能为在同一个包内定义的类型的接收者声明方法
不能为其他包（包括int之类的内建类型）声明方法
*/
import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
