package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "C:\\Users\\Zero\\GolandProjects\\awesomeProject\\src\\file\\text.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error!", err)
		return
	}
	defer file.Close()
	str := "hello\n"
	//写入时，使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//将写入缓冲的数据落盘
	writer.Flush()
}
