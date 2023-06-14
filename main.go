package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("bookstore/view/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("bookstore/view/pages"))))
	//首页
	http.HandleFunc("/index", controller.GetPageBookByPrice)

	//登录
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("/logout", controller.Logout)

	//注册
	http.HandleFunc("/register", controller.Register)
	//获取username 判断用户名是否存在
	http.HandleFunc("/CheckUserName", controller.CheckUserName)
	//http.HandleFunc("/getBook", controller.Getbook)

	//获取分页图书
	http.HandleFunc("/getPageBook", controller.GetPageBook)

	//http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//根据id获取图书，跳转新增或者更新图书界面
	http.HandleFunc("/updateByIDToAddBook", controller.UpdateByIDToAddBook)
	//更新or新增图书
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	//新增图书到购物车
	http.HandleFunc("/addBookCart", controller.AddBookCart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//DeleteCartItemByCartID  删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItemByID)
	//updateCartItem  更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	//去结账
	http.HandleFunc("/checkout", controller.Checkout)

	http.ListenAndServe("localhost:8080", nil)

}
