package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart //判断当前session有没有购物车
	OrderID   string
}
