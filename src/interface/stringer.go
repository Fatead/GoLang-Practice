package main

import "fmt"

// Person 实现了Stringer接口
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	Tom := Person{
		Name: "Tom",
		Age:  12,
	}
	Jim := Person{
		Name: "Jim",
		Age:  14,
	}
	fmt.Println(Tom, Jim)
}
