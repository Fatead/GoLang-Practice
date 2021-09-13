package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	//只有首字母变量大写才能被marshal序列化
	Name   string
	Age    int
	Number string
}

func testStruct() {
	student := Student{
		Name:   "Jack",
		Age:    12,
		Number: "231",
	}
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(data))
}

//将map进行序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "Mick"
	a["age"] = 30
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(data))
}

func main() {
	//testStruct()
	testMap()
}
