package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

//func Getbook(w http.ResponseWriter, r *http.Request) {
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	//获取图书
//	book, err := dao.GetBook()
//	if err != nil {
//		fmt.Println("查询失败")
//		return
//	}
//	files, err := template.ParseFiles(wd + "/bookstore/view/pages/manager/book_manager.html")
//	if err != nil {
//		fmt.Println("跳转失败 err", err)
//		return
//	}
//	files.Execute(w, book)
//}

//添加图书
//func AddBook(w http.ResponseWriter, r *http.Request) {
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
//	book := &model.Book{
//		Title:   title,
//		Author:  author,
//		Price:   fprice,
//		Sales:   int(fsales),
//		Stock:   int(fstock),
//		ImgPath: "/bookstore/view/static/img/logo.gif",
//	}
//	err := dao.AddBook(book)
//	if err != nil {
//		fmt.Println("addbook err", err)
//		return
//	}
//	//调用GetBook 再查询一次数据库
//	Getbook(w, r)
//}

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	err := dao.DeleteBook(bookID)
	if err != nil {
		fmt.Println("删除 err", err)
		return
	}
	//调用GetBook 再查询一次数据库
	GetPageBook(w, r)
}

// 根据id获取图书 跳转到更新或者新增图书界面
func UpdateByIDToAddBook(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	bookID := r.FormValue("bookId")
	book, err := dao.UpdateByIdBook(bookID)
	if err != nil {
		return
	}
	if book.ID > 0 {
		//更新
		files, err := template.ParseFiles(wd + "/bookstore/view/pages/manager/book_add.html")
		if err != nil {
			return
		}
		files.Execute(w, book)
	} else {
		//添加图书
		files, err := template.ParseFiles(wd + "/bookstore/view/pages/manager/book_add.html")
		if err != nil {
			return
		}
		files.Execute(w, "")
	}
}

// 更新或添加图书图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("bookID")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
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
	GetPageBook(w, r)
}

// 获取分页图书
func GetPageBook(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//获取图书
	page, err := dao.GetPageBook(pageNo)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	files, err := template.ParseFiles(wd + "/bookstore/view/pages/manager/book_manager.html")
	if err != nil {
		fmt.Println("跳转失败 err", err)
		return
	}
	files.Execute(w, page)
}

// 获取带分页和价格
func GetPageBookByPrice(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if minPrice == "" && maxPrice == "" {
		page, err = dao.GetPageBook(pageNo)
		if err != nil {
			fmt.Println("查询失败")
			return
		}
	} else {
		page, err = dao.GetPageBookByPrice(pageNo, minPrice, maxPrice)
		//获取图书
		if err != nil {
			fmt.Println("查询失败")
			return
		}
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//获取cookie
	//cookie, err := r.Cookie("user")
	//fmt.Println("cookies :", cookie)
	//调用IsLogin
	faly, session := dao.IsLogin(r)
	if faly {
		//设置page中IsLogin和username字段
		page.IsLogin = true
		page.Username = session.UserName
	}
	files, err := template.ParseFiles(wd + "/bookstore/view/index.html")
	if err != nil {
		fmt.Println("跳转失败 err", err)
		return
	}
	files.Execute(w, page)
}
