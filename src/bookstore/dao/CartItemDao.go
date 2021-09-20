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
func FindCardItemById(bookId int, cardId string) (*model.CartItem, error) {
	sqlStr := "select id, COUNT, amount, book_id, cart_id from cart_items where cart_id = ? and book_id = ?"
	cartItem := &model.CartItem{}
	err := utils.Db.QueryRow(sqlStr, cardId, bookId).Scan(&cartItem.CartItemId, cartItem.Count, cartItem.Amount, &bookId, &cartItem.CartId)
	if err != nil {
		fmt.Println("查询购物项出现错误，没有找到该购物项", err)
		return nil, err
	}
	book, err := FindBookById(bookId)
	if err != nil {
		fmt.Println("查询图书信息出现异常", err)
		return nil, err
	}
	cartItem.Book = book
	return cartItem, nil
}

//findCartItemsByCardId 根据购物车的ID获取购物车中所有的购物项
func findCartItemsByCardId(cardId string) ([]*model.CartItem, error) {
	sqlStr := "select id ,COUNT,amount,book_id,cart_id from cart_itmes where cart_id = ?"
	var cartItems []*model.CartItem
	rows, err := utils.Db.Query(sqlStr, cartItems)
	if err != nil {
		fmt.Println("查询购物车中的购物项过程中出现问题")
		return nil, err
	}
	for rows.Next() {
		cartItem := &model.CartItem{}
		var bookId int
		err := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		if err != nil {
			fmt.Println("扫描写入过程中出现异常")
			return nil, err
		}
		book, err := FindBookById(bookId)
		if err != nil {
			fmt.Println("查询图书信息出现异常，err:", err)
		}
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//UpdateBookCount 更新购物项中图书的数量
func UpdateBookCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_itmes set COUNT = ? , amount = ? where book_id = ? and cart_id = ?"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:", err)
		return err
	}
	//购物车项总额通过结构体方法 GetAmount() 获取
	_, errExec := stmt.Exec(cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return errExec
	}
	//执行成功！
	return nil
}

func DeleteCartItemByCardId() {

}

//DeleteCartItemByCartItemId 根据购物项ID删除购物项
func DeleteCartItemByCartItemId() {

}
