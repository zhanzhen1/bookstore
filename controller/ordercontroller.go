package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// 结账功能
func Checkout(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//获取session
	_, session := dao.IsLogin(r)
	userID := session.UserID
	//获取购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		return
	}
	//生成uuid作为订单号
	uuid := utils.CreateUUID()
	//创建order
	order := &model.Order{
		OrderID:     uuid,
		CreateTime:  time.Now(),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	//将订单保存到数据库中
	dao.AddOrder(order)
	//保存订单项
	//获取购物车中的购物项
	cartItem := cart.CartItem
	for _, item := range cartItem {
		orderItem := &model.OrderItem{
			Count:   item.Count,
			Amount:  item.Amount,
			Title:   item.Book.Title,
			Author:  item.Book.Author,
			Price:   item.Book.Price,
			ImgPath: item.Book.ImgPath,
			OrderID: uuid,
		}
		dao.AddOrderItem(orderItem)
		//更新当前购物车中图书的库存和销量
		book := item.Book
		book.Sales = book.Sales + int(item.Count)
		book.Stock = book.Stock - int(item.Count)
		//更新图书的信息
		dao.UpdateBook(book)
	}
	//清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	//将订单号显示到session中
	session.OrderID = uuid
	//解析模板
	files, err := template.ParseFiles(wd + "/bookstore/view/pages/cart/checkout.html")
	if err != nil {
		return
	}
	files.Execute(w, session)
}
