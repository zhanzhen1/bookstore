package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("查询所有图书", TestGetBook)
	t.Run("测试添加图书", TestAddBook)
}
func TestGetBook(t *testing.T) {
	book, _ := GetBook()
	for i, m := range book {
		fmt.Printf("第%d个图书的信息是%v\n", i+1, m)
	}
}
func TestAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "从入门到入土",
		Author:  "zz",
		Price:   1000000.0,
		Sales:   1000,
		Stock:   1000,
		ImgPath: "/bookstore/view/static/img/logo.gif",
	}
	err := AddBook(book)
	if err != nil {
		fmt.Println("AddBook err", err)
		return
	}
}
func TestDeleteBook(t *testing.T) {
	err := DeleteBook("33")
	if err != nil {
		fmt.Println("delete err", err)
		return
	}
}
func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      34,
		Title:   "从入门到21",
		Author:  "zz",
		Price:   1000.0,
		Sales:   1000,
		Stock:   1000,
		ImgPath: "/bookstore/view/static/img/logo.gif",
	}
	err := UpdateBook(book)
	if err != nil {
		fmt.Println("updateerr", err)
		return
	}
}

func TestGetPageBook(t *testing.T) {
	page, _ := GetPageBook("1")
	fmt.Println("当前页", page.PageNo)
	fmt.Println("总页数", page.TotalPageNo)
	fmt.Println("总记录", page.TotalRecord)
	for i, book := range page.Book {
		fmt.Printf("第%d图书有%v\n", i+1, book)
	}
}
func TestGetPageBookByPrice(t *testing.T) {
	page, _ := GetPageBookByPrice("1", "10", "20")
	fmt.Println("当前页", page.PageNo)
	fmt.Println("总页数", page.TotalPageNo)
	fmt.Println("总记录", page.TotalRecord)
	for i, book := range page.Book {
		fmt.Printf("第%d图书有%v\n", i+1, book)
	}
}
