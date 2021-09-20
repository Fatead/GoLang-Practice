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

//FindCartByUserId 根据ID查找对应的购物车
func FindCartByUserId() {

}

//UpdateCart 更新购物车中的信息
func UpdateCart() {

}

//DeleteCartById 删除购物车中的信息
func DeleteCartById() {

}
