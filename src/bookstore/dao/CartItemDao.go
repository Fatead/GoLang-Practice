package dao

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
)

func AddCartItem(item *model.CartItem) error {
	sqlStr := "insert into cart_items (COUNT,amount,book_id,cart_id) values (?,?,?,?)"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常")
		return err
	}
	_, err = stmt.Exec(&item.Count, &item.Amount, &item.Book.ID, &item.CartId)
	if err != nil {
		fmt.Println("语句执行出错")
		return err
	}
	return nil
}

//FindCardItemById 根据图书的ID和购物车的ID获取对应的购物项
func FindCardItemById() {

}

//findCartItemsByCardId 根据购物车的ID获取购物车中所有的购物项
func findCartItemsByCardId() {

}
