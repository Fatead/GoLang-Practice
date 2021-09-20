package dao

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
)

//FindAllBooks 查询出数据库中已存在的所有图书
func FindAllBooks() ([]*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books"
	query, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询数据出错")
		return nil, err
	}
	var books []*model.Book
	for query.Next() {
		book := &model.Book{}
		err := query.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
		if err != nil {
			fmt.Println("赋值过程中出现异常")
			return nil, err
		}
		books = append(books, book)
	}
	return books, err
}

//FindBookByTitle 根据图书名称查询一条记录
func FindBookByTitle(title string) (book *model.Book, err error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books where title = ?"
	book = &model.Book{}
	err = utils.Db.QueryRow(sqlStr, title).Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
	if err != nil {
		fmt.Println("查询图书信息出错")
		return nil, err
	}
	return book, nil
}

//FindBookById 根据ID查找一本图书
func FindBookById(id int) (book *model.Book, err error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books where id = ?"
	book = &model.Book{}
	err = utils.Db.QueryRow(sqlStr, id).Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
	if err != nil {
		fmt.Println("查询图书信息出错")
		return nil, err
	}
	return book, nil
}

//AddBook 新增一本图书
func AddBook(book *model.Book) error {
	sqlStr := "insert into books (title, author, price, sales, stock, img_path) values (?, ?, ?, ?, ?, ?)"
	prepare, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("新增图书预编译阶段发生错误", err)
		return err
	}
	exec, err := prepare.Exec(&book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
	if exec != nil {
		fmt.Println("新增图书语句执行出错", err)
		return err
	}
	//新增图书成功
	return nil
}

//DeleteBookById 根据图书的ID删除一本书
func DeleteBookById(BookId int) (int64, error) {
	strSql := "delete from books where id = ?"
	prepare, err := utils.Db.Prepare(strSql)
	if err != nil {
		fmt.Println("预编译出错")
		return 0, err
	}
	res, errExec := prepare.Exec(BookId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return 0, errExec
	}
	affect, errRes := res.RowsAffected()
	if errRes != nil {
		fmt.Println("取出受影响行数时出现异常，err:", errRes)
		return 0, errRes
	}
	return affect, nil
}

//updateBooks 更新图书信息, 根据图书的ID对图书信息进行更新
func UpdateBooks(book *model.Book) (int64, error) {
	sqlStr := "Update books set title = ?,author  = ?,price  = ?,sales  = ?,stock  = ?,img_path  = ? where id = ?"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译 Prepare 出错，err :", err)
		return 0, err
	}
	res, errExec := stmt.Exec(&book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath, &book.ID)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return 0, errExec
	}
	affect, errRes := res.RowsAffected()
	if errRes != nil {
		fmt.Println("取出受影响行数时出现异常，err:", errRes)
		return 0, errRes
	}
	return affect, nil
}

//getPage 获取图书总页数 , args是变长参数用来保存sql语句执行中的各种参数信息
func getPage(sqlStr string, pageSize int64, args ...interface{}) (pages int64, count int64, err error) {
	row := utils.Db.QueryRow(sqlStr, args...)
	//得到总数
	err = row.Scan(&count)
	if err != nil {
		fmt.Println("扫描总记录数时出现异常")
		return 0, 0, err
	}
	//计算总页数
	if count%pageSize == 0 {
		pages = count / pageSize
	} else {
		pages = count/pageSize + 1
	}
	return
}

//getBooksForPage 为查询到的图书集合赋值，方便进行分页查询
func getBooksForPage(sqlStr string, args ...interface{}) ([]*model.Book, error) {
	rows, err := utils.Db.Query(sqlStr, args...)
	if err != nil {
		fmt.Println("分页获取出现异常", err)
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
		if err != nil {
			fmt.Println("赋值结果过程中出现异常")
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

//GetPageBooks 返回的Page包含页码，该页包含图书等信息
func GetPageBooks(IndexPage int64) (*model.Page, error) {
	//查询图书的总记录数
	sqlStr := "select count(*) from books"
	var pageSize int64 = 4
	//先获取图书的总页数
	pages, count, err := getPage(sqlStr, pageSize)
	if err != nil {
		fmt.Println("获取页码过程中出现异常", err)
		return nil, err
	}
	//进行分页查询 使用limit关键字来限制返回结果的条数，limit [位置偏移量], 返回数据的数量
	sqlStr2 := "select * from books limit ?, ?"
	books, err := getBooksForPage(sqlStr2, (IndexPage-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println("获取图书集过程中出现异常", err)
		return nil, err
	}
	page := &model.Page{
		Books:     books,
		Pages:     pages,
		PageSize:  pageSize,
		IndexPage: IndexPage,
		Count:     count,
	}
	return page, nil
}

//QueryPrice 查询价格范围内的图书, 将符合条件的图书封装到一个page里面
func QueryPrice(IndexPage int64, low float64, high float64) (*model.Page, error) {
	sqlStr := "select count(*) from books where price between ? and ?"
	var pageSize int64 = 4

	pages, count, err := getPage(sqlStr, pageSize, low, high)
	if err != nil {
		fmt.Println("获取页数页码等出现异常，err:", err)
		return nil, err
	}

	sqlStr2 := "select id , title,author ,price ,sales ,stock ,img_path  from books where price between ? and ? limit ? ,?"
	books, err := getBooksForPage(sqlStr2, low, high, (IndexPage-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println("获取图书集出现异常，err:", err)
		return nil, err
	}

	page := &model.Page{
		Books:     books,
		Pages:     pages,
		PageSize:  pageSize,
		IndexPage: IndexPage,
		Count:     count,
		MinPrice:  low,
		MaxPrice:  high,
	}
	return page, nil
}
