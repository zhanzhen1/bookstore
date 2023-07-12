package main

import (
	"bookstore/controller"
	"bookstore/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	// 加载 .env 配置
	if err := godotenv.Load(); err != nil {
		log.Fatal("loading .env file failed")
	}
	// 初始化数据库
	if err := utils.InitDB(); err != nil {
		log.Fatalf("init database failed: %s", err.Error())
	}
	//创建一个路由
	ginServer := gin.Default()

	// 加载所有 html 模板文件
	files, err := GetDirHTMLFiles("./view")
	if err != nil {
		log.Fatalf("load view folder html files failed: %s", err.Error())
	}
	for _, file := range files {
		fmt.Println(file)
	}
	ginServer.LoadHTMLFiles(files...)
	//加载静态资源
	ginServer.Static("/static", "./static")
	// 重定向到 /index 路由
	ginServer.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/index")
	})
	// 404 页面
	ginServer.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	ginServer.GET("/index", controller.GetPageBookByPrice())
	//跳转图书管理
	ginServer.GET("/getPageBook", controller.GetPageBook())
	//跳转到后台管理
	ginServer.GET("/getManager", controller.GetManager())
	//跳转到登录页面
	ginServer.GET("/getLogin", controller.GetLogin())
	//跳转注册
	ginServer.GET("/getRegister", controller.GetRegister())
	//跳转购物车
	ginServer.GET("/getCart", controller.GetCart())
	//登录
	ginServer.POST("/login", controller.Login())
	//注销
	ginServer.GET("/logout", controller.Logout())
	//根据id跳转新增图书页面or更新图书页面
	ginServer.GET("/updateByIDToAddBook", controller.UpdateByIDToAddBook())
	//更新or新增图书
	ginServer.POST("/updateOrAddBook", controller.UpdateOrAddBook())
	//删除图书
	ginServer.GET("/deleteBook", controller.DeleteBook())
	//注册用户 Register
	ginServer.POST("/register", controller.Register())
	//添加图书到购物车
	ginServer.POST("/addBookCart", controller.AddBookCart())
	//根据用户id获取用户的购物车
	ginServer.GET("/getCartInfo", controller.GetCartInfo())
	//清空购物车
	ginServer.GET("/deleteCart", controller.DeleteCart())
	//删除购物项
	ginServer.GET("/deleteCartItemByID", controller.DeleteCartItemByID())
	//更新购物项
	ginServer.GET("/updateCartItem", controller.UpdateCartItem())
	//结账
	ginServer.GET("/checkout", controller.Checkout())
	////获取username 判断用户名是否存在
	ginServer.POST("/CheckUserName", controller.CheckUserName())
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

// GetDirHTMLFiles get all the html files in the folder
func GetDirHTMLFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, "html") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
