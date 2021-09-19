package model

import "time"

//订单结构体
type Order struct {
	OrderId     string
	CreateTime  time.Time //订单的创建时间
	TotalCount  int64     //订单中商品的总数
	TotalAmount float64   //总钱数
	OrderItems  []*OrderItem
	State       int64 //订单的状态信息： 0->未发货 1->已发货 2->交易完成 -1 -> 取消（退款退货）
	UserId      int
}

func (o *Order) SendComplete() bool {
	return o.State == 1
}

func (o *Order) NoSend() bool {
	return o.State == 0
}

func (o *Order) Complete() bool {
	return o.State == 2
}

func (o *Order) Cancel() bool {
	return o.State == -1
}
