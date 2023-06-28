package main

import (
	"bookstore/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//创建一个路由
	ginServer := gin.Default()
	//全局加载
	//ginServer.LoadHTMLGlob("view/***/**/*")
	//单个文件加载
	//ginServer.LoadHTMLFiles("view/index.html", "view/pages/manager/manager.html",
	//	"view/pages/manager/book_manager.html")
	//加载以.html为后缀的文件
	var files []string
	filepath.Walk("./view", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	ginServer.LoadHTMLFiles(files...)
	//加载静态资源
	ginServer.Static("/static", "./static")
	ginServer.GET("/index", controller.GetPageBookByPrice())
	//跳转图书管理
	ginServer.GET("/getPageBook", controller.GetPageBook())
	//跳转到后台管理
	ginServer.GET("/getManager", controller.GetManager())
	//跳转到登录页面
	ginServer.GET("/getLogin", controller.GetLogin())
	//跳转注册
	ginServer.GET("/getRegister", controller.GetRegister())
	//跳转注册
	ginServer.GET("/getCart", controller.GetCart())
	//登录
	ginServer.POST("/login", controller.Login())
	//注销
	ginServer.GET("/logout", controller.Logout())

	ginServer.Run(":8080")

	//处理静态资源
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static"))))
	//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("view/pages"))))
	//首页
	//http.HandleFunc("/index", controller.GetPageBookByPrice)

	//登录
	//http.HandleFunc("/login", controller.Login)
	////注销
	//http.HandleFunc("/logout", controller.Logout)
	//
	////注册
	//http.HandleFunc("/register", controller.Register)
	////获取username 判断用户名是否存在
	//http.HandleFunc("/CheckUserName", controller.CheckUserName)
	////http.HandleFunc("/getBook", controller.Getbook)
	//
	////获取分页图书
	////http.HandleFunc("/getPageBook", controller.GetPageBook)
	//
	////http.HandleFunc("/addBook", controller.AddBook)
	////删除图书
	//http.HandleFunc("/deleteBook", controller.DeleteBook)
	////根据id获取图书，跳转新增或者更新图书界面
	//http.HandleFunc("/updateByIDToAddBook", controller.UpdateByIDToAddBook)
	////更新or新增图书
	//http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	////新增图书到购物车
	//http.HandleFunc("/addBookCart", controller.AddBookCart)
	////获取购物车信息
	//http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	////清空购物车
	//http.HandleFunc("/deleteCart", controller.DeleteCart)
	////DeleteCartItemByCartID  删除购物项
	//http.HandleFunc("/deleteCartItem", controller.DeleteCartItemByID)
	////updateCartItem  更新购物项
	//http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	////去结账
	//http.HandleFunc("/checkout", controller.Checkout)

	//http.ListenAndServe("localhost:8080", nil)

}
