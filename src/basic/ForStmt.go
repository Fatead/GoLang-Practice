package main

import "fmt"

func main() {

	sum := 0
	// Go的for语句后面的三个构成部分外没有小括号，但是大括号{}是必须的
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//for是Go中的while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

}
