package main

import "fmt"

// 类型[n]T表示拥有n个T类型的值的数组

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	nums := [6]int{1, 2, 3, 4, 5, 6}
	//按照切片进行输出
	fmt.Println(nums[1:4])
}
