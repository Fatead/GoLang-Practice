package controller

import (
	"awesomeProject/src/bookstore/dao"
	"awesomeProject/src/bookstore/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func FindAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dao.FindAllBooks()
	if err != nil {
		fmt.Println("查询图书出现异常")
	}
	t := template.Must(template.ParseFiles("src/bookstore/views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

//GetPageBooks 分页对图书进行查找
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//先从HTTP报文中找到当前用户所在的页
	pageNo := r.PostFormValue("PageNo")
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	//得到页号后根据页号查当前页的图书
	pages, err := dao.GetPageBooks(indexPage)
	if err != nil {
		fmt.Println("查询图书过程中出现错误")
	}
	t := template.Must(template.ParseFiles("src/bookstore/views/pages/manager/book_manager.html"))
	t.Execute(w, pages)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("bookId")
	//将字符串转int类型
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("类型转换出错")
	}
	byId, err := dao.DeleteBookById(id)
	if err != nil {
		fmt.Println("删除图书过程中出现异常", err)
	} else {
		fmt.Println("根据ID删除图书成功，影响到的行数为:", byId)
		// 查找所有图书并跳转到图书管理页面
		GetPageBooks(w, r)
	}
}

//ToUpdateBookPage 跳转到编辑页面，对图书信息进行编辑
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PostFormValue("bookId")
	t := template.Must(template.ParseFiles("src/bookstore/views/pages/manager/book_edit.html"))
	if idStr == "" {
		//新增图书
		t.Execute(w, "")
	} else {
		//对现有的图书信息进行更新
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("解析ID过程中出现异常")
		} else {
			book, _ := dao.FindBookById(id)
			t.Execute(w, book)
		}
	}
}

//AddOrUpdateBook 新增或者更新图书
func AddOrUpdateBook(w http.ResponseWriter, r *http.Request) {
	//对HTTP报文进行解析，得到表单图书的参数
	id, _ := strconv.Atoi(r.PostFormValue("bookId"))
	title := r.PostFormValue("title")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	author := r.PostFormValue("author")
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	book := &model.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
		Sales:  int(sales),
		Stock:  int(stock),
	}
	if id == 0 {
		err := dao.AddBook(book)
		if err != nil {
			fmt.Println("添加图书过程中出现异常")
		}
	} else {
		affect, err := dao.UpdateBooks(book)
		if err != nil {
			fmt.Println("更新图书出现异常")
		} else {
			fmt.Println("更新图书成功，更新影响到的行数为：", affect)
		}
	}
	//跳转到图书管理页面
	GetPageBooks(w, r)
}

func QueryPrice(w http.ResponseWriter, r *http.Request) {
	min := r.FormValue("min")
	max := r.FormValue("max")
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	var low float64
	var high float64
	low, _ = strconv.ParseFloat(min, 64)
	high, _ = strconv.ParseFloat(max, 64)
	pages, err := dao.QueryPrice(indexPage, low, high)
	if err != nil {
		fmt.Println("指定价格范围查询图书出现异常", err)
	}
	isLogin, session, err := dao.IsLogin(r)
	if !isLogin || session == nil {
		fmt.Println("数据库中没查找到该session相关记录，err", err)
		pages.IsLogin = false
	} else {
		pages.IsLogin = true
		pages.UserName = session.Username
	}
	pages.MinPrice = low
	pages.MaxPrice = high
	t := template.Must(template.ParseFiles("src/bookstore/views/searchOfPrice.html"))
	t.Execute(w, pages)
}
