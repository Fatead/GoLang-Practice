package model

// Cart /* 购物车中可以有多种书籍，每种书籍都有一个数量
type Cart struct {
	CartId      string
	TotalCount  int64
	TotalAmount float64
	UserId      int
	CartItem    []*CartItem
}

// GetTotalCount 获得购物车中商品的总数
func (c *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range c.CartItem {
		totalCount += v.Count
	}
	c.TotalCount = totalCount
	return totalCount
}

//GetTotalAmount 获得购物车中商品的总价
func (c *Cart) GetTotalAmount() float64 {
	var amount float64
	for _, v := range c.CartItem {
		amount += v.GetAmount()
	}
	c.TotalAmount = amount
	return amount
}
