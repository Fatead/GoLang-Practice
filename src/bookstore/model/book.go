package model

type Book struct {
	ID        int
	Title     string  //书名
	Author    string  //作者信息
	Price     float64 //书籍价格
	Sales     int     //销量
	Store     int     //库存
	ImagePath string  //书籍图片对应的路径地址

}
