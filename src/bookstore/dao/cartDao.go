package dao

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
)

//添加购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts (id, total_amount, total_count, user_id) values (?, ?, ?, ?)"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("语句预编译过程中出现异常")
		return err
	}
	_, err = stmt.Exec(cart.CartId, cart.TotalAmount, cart.TotalCount, cart.UserId)
	if err != nil {
		fmt.Println("语句执行出错", err)
		return err
	}
	fmt.Println("将数据添加到购物车中成功，开始添加购物车中的数据项")
	for _, v := range cart.CartItems {
		fmt.Println(v)
		// todo : 添加购物车中的数据项到数据库中
	}
	return nil
}

//FindCartByUserId 根据用户ID查找对应的购物车
func FindCartByUserId(userId int) (*model.Cart, error) {
	//先根据用户ID查找其对应的购物车
	strSql := "select id, total_amount,total_count,user_id from carts where user_id = ?"
	cart := &model.Cart{}
	err := utils.Db.QueryRow(strSql, userId).Scan(&cart.CartId, &cart.TotalAmount, &cart.TotalCount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	//然后根据购物车ID查找购物车中的所有商品信息
	items, err := findCartItemsByCardId(cart.CartId)
	if err != nil {
		fmt.Println("查找items出现错误")
	}
	for _, cartItem := range cart.CartItems {
		items = append(items, cartItem)
	}
	cart.CartItems = items
	return cart, nil
}

//UpdateCart 更新购物车中的信息
func UpdateCart(cart *model.Cart) error {
	sql := "update carts set total_amount = ?, total_count = ? where id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("语句预编译过程中发生错误")
		return err
	}
	_, err = stmt.Exec(&cart.TotalAmount, &cart.TotalCount, &cart.CartId)
	if err != nil {
		fmt.Println("更新数据过程中出错")
		return err
	}
	return nil
}

//DeleteCartById 删除购物车中的信息
func DeleteCartById(cartId string) error {
	sqlStr := "delete from carts where id = ?"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("语句预编译过程中发生错误")
		return err
	}
	_, err = stmt.Exec(&cartId)
	if err != nil {
		fmt.Println("执行语句出错")
		return err
	}
	return nil
}
