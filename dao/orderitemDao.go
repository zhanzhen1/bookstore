package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_item(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		fmt.Println("AddOrderItem() err", err)
		return err
	}
	return nil
}
