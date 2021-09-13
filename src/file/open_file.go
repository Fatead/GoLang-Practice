package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\Zero\\GolandProjects\\awesomeProject\\src\\file\\text.txt")
	if err != nil {
		fmt.Println("open file error!", err)
	}
	fmt.Println(file)
	closeErr := file.Close()
	if closeErr != nil {
		fmt.Println("close file error!", closeErr)
	}
}
