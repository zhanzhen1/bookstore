package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

// 新增购物项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "INSERT into cart_item  (count,amount,book_id,cart_id) values (?,?,?,?)"

	_, err := utils.DB.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		fmt.Println("add cartItem err", err)
		return err
	}
	return nil
}

// 根据图书id查询对应的购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select cart_item_id,count,amount, cart_id from cart_item where book_id = ? and cart_id =?"
	row := utils.DB.QueryRow(sqlStr, bookID, cartID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		fmt.Println("scan bookId err", err)
		return nil, err
	}
	//根据bookId查询图书信息
	book, err := GetBookByID(bookID)
	if err != nil {
		return nil, err
	}
	cartItem.Book = book
	return cartItem, nil
}

// 根据购物车的id获取购物车中所有的购物项
func GetCartItemByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select cart_item_id,count,amount,book_id,cart_id from cart_item where cart_id = ?"
	stmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		return nil, nil
	}
	rows, err := stmt.Query(cartID)
	if err != nil {
		fmt.Println("query err", err)
		return nil, nil
	}
	//row, err := utils.DB.Query(sqlStr, cartID)
	//if err != nil {
	//	fmt.Println("query err", err)
	//	return nil, err
	//}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err != nil {
			fmt.Println("Scan err", err)
			return nil, nil
		}
		book, err := GetBookByID(bookID)
		if err != nil {
			return nil, err
		}
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

// 更新购物项的图书数量
func UpdateCartItemBookCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_item set count = ? ,amount = ? where book_id = ? and cart_id =?"
	_, err := utils.DB.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		fmt.Println("UpdateCartItemBookCount err", err)
		return err
	}
	return nil
}

// 根据购物车的id删除购物项
func DeleteCartItemByCartID(cartID string) error {
	sqlStr := "delete from cart_item where cart_id = ?"
	_, err := utils.DB.Exec(sqlStr, cartID)
	if err != nil {
		fmt.Println("DeleteCartItemByCartID err", err)
		return err
	}
	return nil
}

// 根据购物项的id删除购物项
func DeleteCartItemByID(cartItemID string) error {
	sqlStr := "delete from cart_item where cart_item_id = ?"
	_, err := utils.DB.Exec(sqlStr, cartItemID)
	if err != nil {
		fmt.Println("DeleteCartItemByID err", err)
		return err
	}
	return nil
}
