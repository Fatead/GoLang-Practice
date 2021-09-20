package test

import (
	"awesomeProject/src/bookstore/dao"
	"awesomeProject/src/bookstore/model"
	"fmt"
	"testing"
)

var book = &model.Book{
	Title:     "Mysql必知必会",
	Author:    "Ben Forta",
	Price:     33,
	Sales:     99,
	Stock:     100,
	ImagePath: "src/bookstore/views/static/img/default.jpg",
}

func TestFindAllBooks(t *testing.T) {
	books, err := dao.FindAllBooks()
	if err != nil {
		fmt.Println("查找所有书籍过程中发生错误")
		return
	}
	fmt.Println("查找所有书籍信息成功，所有书籍信息如下：")
	for _, book := range books {
		fmt.Println(book)
	}
}

func TestFindBookById(t *testing.T) {
	book, err := dao.FindBookById(10)
	if err != nil {
		fmt.Println("根据ID查找书籍信息发生错误")
		return
	}
	fmt.Println("查找书籍信息成功")
	fmt.Println(book)
}

func TestFindBookByTitle(t *testing.T) {
	book, err := dao.FindBookByTitle("小王子")
	if err != nil {
		fmt.Println("根据Title查找书籍信息发生错误")
		return
	}
	fmt.Println("查找书籍信息成功")
	fmt.Println(book)
}

func TestAddBook(t *testing.T) {
	err := dao.AddBook(book)
	if err != nil {
		fmt.Println("添加图书信息失败")
	} else {
		//验证数据库中是否有已经添加的图书信息
		book, err := dao.FindBookByTitle(book.Title)
		if err != nil {
			fmt.Println("数据库中没有该书籍的信息，添加图书信息失败")
		} else {
			fmt.Println("添加图书信息成功，图书信息如下：")
			fmt.Println(book)
		}
	}
}

func TestDeleteBookById(t *testing.T) {
	res, err := dao.DeleteBookById(45)
	if err != nil {
		fmt.Println("删除图书信息失败")
	} else {
		fmt.Println("删除图书信息成功，受影响的行数为：", res)
	}
}

func TestUpdateBooks(t *testing.T) {

}

func TestGetPageBooks(t *testing.T) {

}

func TestQueryPrice(t *testing.T) {

}
