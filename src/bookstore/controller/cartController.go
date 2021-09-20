package controller

import (
	"awesomeProject/src/bookstore/dao"
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
	"net/http"
	"strconv"
)

//AddBook2Cart 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	login, session, _ := dao.IsLogin(r)
	if login {
		bookId, _ := strconv.Atoi(r.PostFormValue("bookId"))
		book, _ := dao.FindBookById(bookId)
		cart, _ := dao.FindCartByUserId(session.UserId)
		if cart == nil {
			//该用户还未创建购物车
			cartId := utils.CreateUUID()
			cartItem := &model.CartItem{
				CartId: cartId,
				Count:  1,
				Book:   book,
			}
			var cartItems []*model.CartItem
			cartItems = append(cartItems, cartItem)
			cart := &model.Cart{
				CartId:    cartId,
				UserId:    session.UserId,
				CartItems: cartItems,
			}
			cart.TotalAmount = cart.GetTotalAmount()
			cart.TotalCount = cart.GetTotalCount()
			err := dao.AddCart(cart)
			if err != nil {
				fmt.Println("新建购物车失败")
			}
		} else {
			//该用户已经有购物车，需要检查是否已经创建购物项
			cartItem, _ := dao.FindCardItemById(bookId, cart.CartId)
			if cartItem == nil {
				//需要新创建购物项
				cartId := utils.CreateUUID()
				cartItem := &model.CartItem{
					CartId: cartId,
					Count:  1,
					Book:   book,
					Amount: book.Price,
				}
				var cartItems []*model.CartItem
				cartItems = append(cartItems, cartItem)
				cart.CartItems = cartItems
			} else {
				//已经创建该购物项，则购物车内购物项数量加1

			}
		}
	} else {
		w.Write([]byte("请先登录再进行操作"))
	}
}

//GetCartInfo 获取用户当前的购物车详情
func GetCartInfo(w http.ResponseWriter, r *http.Request) {

}

//DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {

}

//DeleteCartItem 根据ID删除购物车中的商品
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {

}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {

}
