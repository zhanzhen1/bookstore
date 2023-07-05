package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 定义全局page
var page *model.Page
var err error

// 删除图书
//func DeleteBook1(w http.ResponseWriter, r *http.Request) {
//	bookID := r.FormValue("bookId")
//	err := dao.DeleteBook(bookID)
//	if err != nil {
//		fmt.Println("删除 err", err)
//		return
//	}
//	//调用GetBook 再查询一次数据库
//	GetPageBook()
//}

// gin 优化删除图书
func DeleteBook() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		bookID := ctx.Query("bookId")
		err = dao.DeleteBook(bookID)
		if err != nil {
			fmt.Println("删除 err", err)
			return
		}
		//调用GetBook 再查询一次数据库
		page, err = dao.GetPageBook(GetPageNo(ctx))
		if err != nil {
			ReturnErr(ctx, err)
			return
		}
		ctx.HTML(http.StatusOK, "book_manager.html", page)
	}

}

// 根据id获取图书 跳转到更新或者新增图书界面
//func UpdateByIDToAddBook1(w http.ResponseWriter, r *http.Request) {
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	bookID := r.FormValue("bookId")
//	book, err := dao.UpdateByIdBook(bookID)
//	if err != nil {
//		return
//	}
//	if book.ID > 0 {
//		//更新
//		files, err := template.ParseFiles(wd + "/bookstore/view/pages/manager/book_add.html")
//		if err != nil {
//			return
//		}
//		files.Execute(w, book)
//	} else {
//		//添加图书
//		files, err := template.ParseFiles(wd + "/bookstore/view/pages/manager/book_add.html")
//		if err != nil {
//			return
//		}
//		files.Execute(w, "")
//	}
//}

// gin 优化根据id获取图书 跳转到更新或者新增图书界面
func UpdateByIDToAddBook() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		bookID := context.Query("bookId")
		book, err := dao.UpdateByIdBook(bookID)
		if err != nil {
			return
		}
		if book.ID > 0 {
			//更新
			context.HTML(http.StatusOK, "book_add.html", book)
		} else {
			//添加图书
			context.HTML(http.StatusOK, "book_add.html", nil)
		}
	}

}

// 更新或添加图书图书
//func UpdateOrAddBook1(w http.ResponseWriter, r *http.Request) {
//	bookID := r.PostFormValue("bookID")
//	title := r.PostFormValue("title")
//	price := r.PostFormValue("price")
//	author := r.PostFormValue("author")
//	sales := r.PostFormValue("sales")
//	stock := r.PostFormValue("stock")
//	//将价格和库存进行转换
//	fprice, _ := strconv.ParseFloat(price, 64)
//	//base 进制 bitsize 类型
//	fsales, _ := strconv.ParseInt(sales, 10, 0)
//	fstock, _ := strconv.ParseInt(stock, 10, 0)
//	fbookID, _ := strconv.ParseInt(bookID, 10, 0)
//	book := &model.Book{
//		ID:      int(fbookID),
//		Title:   title,
//		Author:  author,
//		Price:   fprice,
//		Sales:   int(fsales),
//		Stock:   int(fstock),
//		ImgPath: "/bookstore/view/static/img/default.jpg",
//	}
//	if book.ID > 0 {
//		//修改
//		err := dao.UpdateBook(book)
//		if err != nil {
//			fmt.Println("update err", err)
//			return
//		}
//	} else {
//		//添加
//		err := dao.AddBook(book)
//		if err != nil {
//			fmt.Println("add err", err)
//			return
//		}
//	}
//	//调用GetBook 再查询一次数据库
//	GetPageBook()
//}

// 更新或添加图书图书
func UpdateOrAddBook() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		bookID := ctx.PostForm("bookID")
		title := ctx.PostForm("title")
		price := ctx.PostForm("price")
		author := ctx.PostForm("author")
		sales := ctx.PostForm("sales")
		stock := ctx.PostForm("stock")
		//将价格和库存进行转换
		fprice, _ := strconv.ParseFloat(price, 64)
		//base 进制 bitsize 类型
		fsales, _ := strconv.ParseInt(sales, 10, 0)
		fstock, _ := strconv.ParseInt(stock, 10, 0)
		fbookID, _ := strconv.ParseInt(bookID, 10, 0)
		book := &model.Book{
			ID:      int(fbookID),
			Title:   title,
			Author:  author,
			Price:   fprice,
			Sales:   int(fsales),
			Stock:   int(fstock),
			ImgPath: "/bookstore/view/static/img/default.jpg",
		}
		if book.ID > 0 {
			//修改
			err := dao.UpdateBook(book)
			if err != nil {
				fmt.Println("update err", err)
				return
			}
		} else {
			//添加
			err := dao.AddBook(book)
			if err != nil {
				fmt.Println("add err", err)
				return
			}
		}
		//调用GetBook 再查询一次数据库
		page, err = dao.GetPageBook(GetPageNo(ctx))
		if err != nil {
			ReturnErr(ctx, err)
			log.Println("查询图书失败", err)
			return
		}
		ctx.HTML(http.StatusOK, "book_manager.html", page)
	}
}

var pageNo string

// 获取分页图书
func GetPageBook() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		//获取图书
		page, err = dao.GetPageBook(GetPageNo(ctx))
		if err != nil {
			fmt.Println("查询失败")
			ReturnErr(ctx, err)
			return
		}
		ctx.HTML(http.StatusOK, "book_manager.html", page)
	}

}

// 获取带分页和价格
//func GetPageBookByPrice1(w http.ResponseWriter, r *http.Request) {
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	pageNo := r.FormValue("pageNo")
//	if pageNo == "" {
//		pageNo = "1"
//	}
//	var page *model.Page
//	minPrice := r.FormValue("min")
//	maxPrice := r.FormValue("max")
//	if minPrice == "" && maxPrice == "" {
//		page, err = dao.GetPageBook(pageNo)
//		if err != nil {
//			fmt.Println("查询失败")
//			return
//		}
//	} else {
//		page, err = dao.GetPageBookByPrice(pageNo, minPrice, maxPrice)
//		//获取图书
//		if err != nil {
//			fmt.Println("查询失败")
//			return
//		}
//		//将价格范围设置到page中
//		page.MinPrice = minPrice
//		page.MaxPrice = maxPrice
//	}
//	//获取cookie
//	//cookie, err := r.Cookie("user")
//	//fmt.Println("cookies :", cookie)
//	//调用IsLogin
//	faly, session := dao.IsLogin()
//	if faly {
//		//设置page中IsLogin和username字段
//		page.IsLogin = true
//		page.Username = session.UserName
//	}
//	files, err := template.ParseFiles(wd + "/view/index.html")
//	if err != nil {
//		fmt.Println("跳转失败 err", err)
//		return
//	}
//	files.Execute(w, page)
//}

// gin index
func GetPageBookByPrice() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		var page *model.Page
		var err error
		minPrice := ctx.Query("min")
		maxPrice := ctx.Query("max")
		pageNo := GetPageNo(ctx)
		//如果minPrice=空值就是翻页
		if minPrice == "" && maxPrice == "" {
			page, err = dao.GetPageBook(pageNo)
			if err != nil {
				fmt.Println("GetPageBook()查询失败", err)
				ReturnErr(ctx, err)
				return
			}
			fmt.Println("minPrice:", minPrice)
		} else { //查询
			page, err = dao.GetPageBookByPrice(pageNo, minPrice, maxPrice)
			//获取图书
			if err != nil {
				fmt.Println("GetPageBookByPrice查询失败", err)
				ReturnErr(ctx, err)
				return
			}
			//将价格范围设置到page中
			page.MinPrice = minPrice
			page.MaxPrice = maxPrice
		}
		//调用IsLogin
		faly, session := dao.IsLogin(ctx)
		if faly {
			//设置page中IsLogin和username字段
			page.IsLogin = true
			page.Username = session.UserName
		}
		ctx.HTML(http.StatusOK, "index.html", page)
	}
}
