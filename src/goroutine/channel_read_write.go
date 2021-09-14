package main

func main() {
	//管道可以声明为制度或者只写
	//在默认情况下，管道是双向的，

	//声明为只写
	var channel_write chan<- int
	channel_write = make(chan int, 2)
	channel_write <- 20

	//声明为只读
	var channel_read <-chan int
	channel_read = make(chan int, 2)
	<-channel_read
}
