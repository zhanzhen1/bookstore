package model

// 订单项
type OrderItem struct {
	OrderItemID int64
	Count       int64   //订单项中图书的数量
	Amount      float64 //订单项中图书的金额
	Title       string  //订单项中图书的书名
	Author      string  //订单项中图书的作者
	Price       float64 //订单项中图书的价格
	ImgPath     string  //订单项中图书的图片
	OrderID     string  //订单项所属的订单
}
