package test

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
	"testing"
)

var book1 = &model.Book{
	ID:    1,
	Price: 27.00,
}

var book2 = &model.Book{
	ID:    2,
	Price: 23.00,
}

func TestAddCart(t *testing.T) {
	//创建一个购物车的切片
	var cartItems []*model.CartItem
	cartUUID := utils.CreateUUID()
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartId: cartUUID,
	}
	cartItems = append(cartItems, cartItem1)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  12,
		CartId: cartUUID,
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartId:    cartUUID,
		CartItems: cartItems,
		UserId:    1,
	}
	fmt.Println(cart)
}

func TestFindCartByUserID(t *testing.T) {

}

func TestUpdateCart(t *testing.T) {

}

func TestDeleteCart(t *testing.T) {

}
