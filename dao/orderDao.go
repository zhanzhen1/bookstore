package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

// 添加订单
func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders(order_id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		fmt.Println("AddOrder() err", err)
		return err
	}
	return nil
}
