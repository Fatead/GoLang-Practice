package model

// OrderItem /*订单中的订单项
type OrderItem struct {
	ID      int
	Count   int64
	Amount  float64
	Title   string
	Author  string
	Price   float64
	ImaPath string
	OrderId string
}
