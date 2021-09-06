package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	*p = 21
	*p = *p + 10
	fmt.Println(*p, j)
}
