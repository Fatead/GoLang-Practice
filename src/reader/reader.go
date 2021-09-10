package main

import (
	"fmt"
	"io"
	"strings"
)

/*
IO包指定了io.Reader接口，表示从数据流的末尾进行读取，Reader接口有一个Read方法
func (T) Read(b []byte) (n int, err error)
Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
*/

func main() {
	r := strings.NewReader("hello reader")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
