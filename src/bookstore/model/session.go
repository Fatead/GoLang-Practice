package model

type Session struct {
	SessionId string
	UserId    int
	Username  string
	Cart      *Cart
	OrderId   string
	Orders    []*Order
}
