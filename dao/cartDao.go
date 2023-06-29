package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

// 新增购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "INSERT INTO cart (cart_id,total_count,total_amount,user_id) VALUES(?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		fmt.Println("add cart err", err)
		return err
	}
	cartItem := cart.CartItem
	for _, item := range cartItem {
		//将购物项保存到数据库中
		AddCartItem(item)
	}
	return nil
}

// 根据用户id从数据库中查询对应的购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select * from cart where user_id = ?"
	stmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(userID)
	cart := &model.Cart{}
	err = row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	fmt.Println(err)
	if err != nil {
		fmt.Println("GetCartByUserID() scan userID err", err)
		return nil, err
	}
	//row := utils.DB.QueryRow(sqlStr, userID)
	//cart := &model.Cart{}
	//err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	//if err != nil {
	//	fmt.Println("scan userID err", err)
	//	return nil, err
	//}
	//获取当前的购物项
	cartItem, _ := GetCartItemByCartID(cart.CartID)
	//将所有的购物项设置到购物车中
	cart.CartItem = cartItem
	return cart, err
}

// 更新购物车中图书的总数量和总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update cart set total_count = ?,total_amount = ? where cart_id = ?"
	_, err := utils.DB.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据购物车id删除购物车
func DeleteCartByCartID(cartID string) error {
	//删除购物车之前需要先删除所有购物项
	err := DeleteCartItemByCartID(cartID)
	if err != nil {
		return err
	}
	sqlStr := "delete from cart where cart_id = ?"
	_, err = utils.DB.Exec(sqlStr, cartID)
	if err != nil {
		fmt.Println("DeleteCartByCartID() err", err)
		return err
	}
	return nil
}
