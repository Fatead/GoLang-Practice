package model

type Page struct {
	Pages     int64 //总页数
	PageSize  int64 //每页显示商品的条数
	Count     int64 //记录数
	IndexPage int64 //当前页
	Books     []*Book
	MinPrice  float64
	MaxPrice  float64
	IsLogin   bool
	UserName  string
}
