package main

import "fmt"

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primechan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
		}
	}
	fmt.Println("有一个primeNum取不到数据，退出了")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exit管道取出了4个结果
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数为%d", res)
	}
}
