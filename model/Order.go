package model

import "time"

// 订单
type Order struct {
	OrderID     string    //用uuid来生成
	CreateTime  time.Time //生成订单时间
	TotalCount  int64     //订单中的总数量
	TotalAmount float64   //订单中的总金额
	State       int64     //订单中的状态 0 未发货 1 已发货
	UserID      int64     //订单所属的用户

}
