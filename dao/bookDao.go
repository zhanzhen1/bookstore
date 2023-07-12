package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

// gorm 获取所有图书
func GetBook() ([]*model.Book, error) {
	var book []*model.Book
	utils.Db1.Find(&book)
	return book, nil
}

// 获取所有图书
//
//	func GetBook1() ([]*model.Book, error) {
//		sqlStr := "select * from book"
//		query, err := utils.DB.Query(sqlStr)
//		if err != nil {
//			fmt.Println("查询err", err)
//			return nil, err
//		}
//		var book []*model.Book
//		for query.Next() {
//			books := &model.Book{}
//			err := query.Scan(&books.ID, &books.Title, &books.Author, &books.Price, &books.Sales, &books.Stock, &books.ImgPath)
//			if err != nil {
//				fmt.Println("query err", err)
//				return nil, err
//			}
//			book = append(book, books)
//		}
//		return book, nil
//	}

// 根据bookid查询
func GetBookByID(id string) (*model.Book, error) {
	sqlStr := "select * from book where id = ?"
	row := utils.DB.QueryRow(sqlStr, id)
	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		fmt.Println("GetBookByID() err", err)
		return nil, err
	}
	return book, nil
}

// 新增图书

func AddBook(book *model.Book) error {
	sqlStr := "insert into book (title,author,price,sales,stock,img_path) values (?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, &book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		fmt.Println("exec err", err)
		return nil
	}
	return err
}

// 根据图书id删除

func DeleteBook(id string) error {
	sqlStr := "DELETE FROM book WHERE id = ?"
	_, err := utils.DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败 err", err)
		return err
	}
	return err
}

// 根据id更新图书内容

func UpdateByIdBook(id string) (*model.Book, error) {
	sqlStr := "select *  FROM book WHERE id = ?"
	row := utils.DB.QueryRow(sqlStr, id)
	book := &model.Book{}
	row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Price,
		&book.Sales,
		&book.Stock,
		&book.ImgPath,
	)
	return book, nil
}

// 更新图书

func UpdateBook(book *model.Book) error {
	sqlStr := "UPDATE book  SET title = ?,author = ? ,price= ? ,sales = ? ,stock = ?  where id = ?"

	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

var book *model.Book
var bookList []*model.Book

// 获取带分页的图书信息
func GetPageBook(pageNo int) (*model.Page, error) {
	//获取数据库图书的总数
	var totalRecord int
	utils.Db1.Model(&model.Book{}).Select("count(*) as count ").Find(&book).Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int
	pageSize = 4
	//获取总页数
	var totalPageNo int
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//sqlStr2 := "select * from book limit ?,?"
	//query2, err := utils.DB.Query(sqlStr2, (pageNo-1)*pageSize, pageSize)
	utils.Db1.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&bookList)

	//for query2.Next() {
	//	book := &model.Book{}
	//	query2.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	//	//添加到切片books中
	//	books = append(books, book)
	//}

	//创建page
	page := &model.Page{
		Book:        bookList,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
		PageNo:      pageNo,
		PageSize:    pageSize,
	}
	return page, nil
}

// 带价格查询
func GetPageBookByPrice(pageNo int, minPrice string, maxPrice string) (*model.Page, error) {
	//获取数据库图书的总数
	var totalRecord int
	utils.Db1.Model(&model.Book{}).Select("count(*) as count").Where("price between ? and ?", minPrice, maxPrice).Find(&book).Scan(&totalRecord)
	//row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int
	pageSize = 4
	//获取总页数
	var totalPageNo int
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	utils.Db1.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&bookList)
	//创建page
	page := &model.Page{
		Book:        bookList,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
		PageNo:      pageNo,
		PageSize:    pageSize,
	}
	return page, nil
}

//// 获取带分页的图书信息
//func GetPageBook1(pageNo int64) (*model.Page, error) {
//	//获取数据库图书的总数
//	sqlStr := "select count(*) from book "
//	var totalRecord int64
//	row := utils.DB.QueryRow(sqlStr)
//	row.Scan(&totalRecord)
//	//设置每页只显示4条记录
//	var pageSize int64
//	pageSize = 4
//	//获取总页数
//	var totalPageNo int64
//	if totalRecord%pageSize == 0 {
//		totalPageNo = totalRecord / pageSize
//	} else {
//		totalPageNo = totalRecord/pageSize + 1
//	}
//	sqlStr2 := "select * from book limit ?,?"
//	query2, err := utils.DB.Query(sqlStr2, (pageNo-1)*pageSize, pageSize)
//	if err != nil {
//		return nil, err
//	}
//	books := []*model.Book{}
//	for query2.Next() {
//		book := &model.Book{}
//		query2.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
//		//添加到切片books中
//		books = append(books, book)
//	}
//	//创建page
//	page := &model.Page{
//		Book:        books,
//		TotalPageNo: totalPageNo,
//		TotalRecord: totalRecord,
//		PageNo:      pageNo,
//		PageSize:    pageSize,
//	}
//	return page, nil
//}
//
//// 带价格查询
//func GetPageBookByPrice1(pageNo int64, minPrice string, maxPrice string) (*model.Page, error) {
//	//获取数据库图书的总数
//	sqlStr := "select count(*) from book where price between ? and	?"
//	var totalRecord int64
//	row := utils.DB.QueryRow(sqlStr, minPrice, maxPrice)
//	row.Scan(&totalRecord)
//	//设置每页只显示4条记录
//	var pageSize int64
//	pageSize = 4
//	//获取总页数
//	var totalPageNo int64
//	if totalRecord%pageSize == 0 {
//		totalPageNo = totalRecord / pageSize
//	} else {
//		totalPageNo = totalRecord/pageSize + 1
//	}
//	sqlStr2 := "select * from book where price between ? and ? limit ?,?"
//	query2, err := utils.DB.Query(sqlStr2, minPrice, maxPrice, (pageNo-1)*pageSize, pageSize)
//	if err != nil {
//		return nil, err
//	}
//	books := []*model.Book{}
//	for query2.Next() {
//		book := &model.Book{}
//		query2.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
//		//添加到切片books中
//		books = append(books, book)
//	}
//	//创建page
//	page := &model.Page{
//		Book:        books,
//		TotalPageNo: totalPageNo,
//		TotalRecord: totalRecord,
//		PageNo:      pageNo,
//		PageSize:    pageSize,
//	}
//	return page, nil
//}
