package main

import (
	"fmt"
	"sort"
)

//利用接口对学生成绩按照降序进行排序

type Student struct {
	name  string
	age   int
	score int
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	var students StudentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			name:  "some",
			age:   20,
			score: 90 + i,
		}
		students = append(students, student)
	}
	sort.Sort(students)
	for _, student := range students {
		fmt.Println(student)
	}

}
