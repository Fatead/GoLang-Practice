package model

type CartItem struct {
	CartItemId int64
	Count      int64
	Amount     float64
	Book       *Book
	CartId     string
}

//获取购物车中某项商品的总价
func (c *CartItem) GetAmount() float64 {
	amount := float64(c.Count) * c.Book.Price
	c.Amount = amount
	return amount
}
