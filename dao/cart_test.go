package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestAddCart(t *testing.T) {
	book := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	var cartitems []*model.CartItem
	//创建两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  2,
		CartID: "66666",
	}
	cartitems = append(cartitems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  3,
		CartID: "66666",
	}
	cartitems = append(cartitems, cartItem2)
	cart := &model.Cart{
		CartID:   "66666",
		CartItem: cartitems,
		UserID:   2,
	}
	//将购物车插入到数据中
	AddCart(cart)
}

func TestGetCartItemByBookID(t *testing.T) {
	bookID, err := GetCartItemByBookIDAndCartID("1", "66666")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(bookID)
}
func TestGetCartItemByCartID(t *testing.T) {
	cartID, err := GetCartItemByCartID("1")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	for i, item := range cartID {
		fmt.Printf("当前%d的购物项是%v\n", i+1, item)
	}
}

func TestGetCartByUserID(t *testing.T) {
	userID, err := GetCartByUserID(3)
	if err != nil {
		return
	}
	fmt.Printf("id为2的购物车信息是：%v", userID)
}

func TestUpdateCartItemBookCount(t *testing.T) {
	//UpdateCartItemBookCount()
}

func TestDeleteCartByCartID(t *testing.T) {
	fmt.Println("TestDeleteCartByCartID() ")
	DeleteCartByCartID("5b0154cc-a09a-40dc-76fb%!(EXTRA []uint8=[255 235 102 229 161 195])")
}
func TestDeleteCartItemByID(t *testing.T) {
	fmt.Println("删除成功")
	err := DeleteCartItemByID("19")
	if err != nil {
		fmt.Println("DeleteCartItemByID err", err)
		return
	}
}
