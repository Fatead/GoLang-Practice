package main

/*
 *   Go的基本类型有:
 *   bool
 *   string
 *   int int8 int16 int32 int64
 *   unit uint8 uint16 uint32 uint64 uintptr
 *   byte uint8的别名
 *   rune int32的别名，表示一个Unicode码点
 *   float32 float64
 *   complex64 complex128 复数
 *   没有明确初始值的变量会被赋予它们的零值：
 *        数值类型的零值为0，
 *        布尔类型为false
 *        字符串为 “” (空字符串)
 */

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
