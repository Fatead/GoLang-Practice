package main

// 函数值，函数也是值，它们可以像其他值一样传递

import (
	"fmt"
	"math"
)

/*
 * 对于以小写字母开头的函数，其作用域只属于所声明的包内，不能被其他包使用
 * 如果我们把函数名以大写字母开头，该函数就可以被其他包调用
 */

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}
