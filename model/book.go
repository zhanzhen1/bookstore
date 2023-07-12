package model

// 图书
type Book struct {
	ID      int
	Title   string
	Author  string
	Price   float64
	Sales   int
	Stock   int
	ImgPath string
	Count   int
}
