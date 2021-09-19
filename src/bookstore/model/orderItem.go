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

func (p *Page) IsHasPre() bool {
	return p.IndexPage > 1
}

func (p *Page) IsHasNext() bool {
	return p.IndexPage < p.Pages
}

func (p *Page) GetPrevPageNo() int64 {
	return p.IndexPage - 1
}

func (p *Page) GetNextPageNo() int64 {
	return p.IndexPage + 1
}
