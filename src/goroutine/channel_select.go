package main

import "fmt"

func main() {

	//使用select可以解决从管道取数据的堵塞问题

	//定义一个管道，10个int数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//定义一个管道，5个string数据
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭而阻塞会导致deadLock
	//在实际开发中我们不好确定什么时候关闭该管道，可以使用select来解决
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan中读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan中读取的数据%v\n", v)
		default:
			fmt.Printf("都取不到数据\n")
			return
		}

	}

}
