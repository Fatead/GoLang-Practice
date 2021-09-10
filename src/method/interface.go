package main

/*
接口是一个方法集，
在Go中，只要结构体的方法实现了（包含了）接口中的所有方法，
就可以说这个结构体实现了这个接口
*/

import "fmt"

/*
 * 定义一个接口，包含write()方法
 */
type act interface {
	write()
}

type xiaoming struct {
	//xiaoming的结构体
}

type xiaofang struct {
}

func (xm *xiaoming) write() {
	//xiaoming的结构体方法write，接收者为指针类型，即xiaoming实现了act接口
	fmt.Println("xiaoming write")
}

func (xf *xiaofang) write() {
	fmt.Println("xiaofang write")
}

func main() {
	var w act
	xm := xiaoming{}
	xf := xiaofang{}
	w = &xm //实例化接口
	w.write()
	w = &xf
	w.write()
}
