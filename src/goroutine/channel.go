package main

import "fmt"

/*
信道
信道是带有类型的管道，可以通过它用信道操作符“<-”来发送或者接收值
ch <- v 将V发送到信道ch
v := <-ch 从ch接收值并赋予v
箭头的方向就是数据流的方向
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

/*
关键字make和new的区别
new分配空间，传递给new函数的参数是一个类型，返回值是指向这个新分配的零值的指针

//为默认数据类型分配空间
var num *int
num = new(int)

//为自定义类型分配空间
type Student struct{
     name String
     age int
}

var stu* Student
stu = new(Student)


make
channel\map\slice 类型是引用类型，make用于分配内存，
但返回的类型就是这三个类型的引用，而不是指针
//声明并用make函数进行初始化，指定长度为3
s := make([]int 3)
*/

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
