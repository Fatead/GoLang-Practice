package te

import (
	"testing"
)

// Go中的测试类样例
func TestArray01(t *testing.T) {
	var a [3]int
	a[0] = 1
	b := [3]int{11, 22, 33}
	c := [2][2]int{
		{11, 22},
		{33, 44},
	}
	for _, i := range b {
		t.Log(i)
	}
	for _, i := range c {
		t.Log(i)
	}
}

/*
切片也称为动态数组，可以灵活的扩容
切片可以使用数组来初始化，也可以通过内置函数make来进行初始化
*/
func TestArraySection(t *testing.T) {
	arr := []int{11, 22, 33, 44, 55}
	arr_sec := arr[:3]
	arr = append(arr, 10)
	t.Log(arr)
	t.Log(arr_sec)
}
