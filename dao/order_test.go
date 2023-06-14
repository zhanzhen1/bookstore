package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"testing"
	"time"
)

func TestAddOrder(t *testing.T) {
	//生成uuid
	uuid := utils.CreateUUID()
	order := &model.Order{
		OrderID:     uuid,
		CreateTime:  time.Now(),
		TotalCount:  2,
		TotalAmount: 200,
		State:       0,
		UserID:      2,
	}
	//创建订单项
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "入门到入土",
		Author:  "zz",
		Price:   100,
		ImgPath: "bookstore/view/static/img/default.jpg",
		OrderID: uuid,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "入门到入土11",
		Author:  "zz",
		Price:   100,
		ImgPath: "bookstore/view/static/img/default.jpg",
		OrderID: uuid,
	}
	//保存订单
	AddOrder(order)
	//保存订单项
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}
