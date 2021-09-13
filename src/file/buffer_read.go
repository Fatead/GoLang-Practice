package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\Zero\\GolandProjects\\awesomeProject\\src\\file\\text.txt")
	if err != nil {
		fmt.Println("open file error!", err)
	}
	defer file.Close()
	//创建一个带缓冲区的Reader
	reader := bufio.NewReader(file)
	for {
		str, err2 := reader.ReadString('\n') //读到一个换行就结束
		if err2 == io.EOF {                  //io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
}
