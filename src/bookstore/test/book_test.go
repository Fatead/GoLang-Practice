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
	book, err := dao.FindBookById(23)
	if err != nil {
		fmt.Println("根据ID查询图书信息出现异常")
	}
	book.Price = 85
	res, err := dao.UpdateBooks(book)
	if err != nil {
		fmt.Println("更新图书信息过程中出现异常")
	} else {
		fmt.Println("图书更新成功，受影响的行数为：", res)
	}
}

func TestGetPageBooks(t *testing.T) {
	page, err := dao.GetPageBooks(2)
	if err != nil || page == nil {
		fmt.Println("分页查询图书出现异常， err", err)
	}
	fmt.Println("查询图书信息成功，图书信息如下：")
	for _, book := range page.Books {
		fmt.Println(book)
	}
}

// 测试根据价格范围查找符合要求的图书
func TestQueryPrice(t *testing.T) {
	pages, err := dao.QueryPrice(2, 30, 40)
	if err != nil {
		fmt.Println("分页查询图书过程中出现异常")
	}
	for _, book := range pages.Books {
		fmt.Println(*book)
	}
}
