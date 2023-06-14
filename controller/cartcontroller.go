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

// 添加图书到购物车
func AddBookCart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		bookID := r.FormValue("bookID")
		fmt.Println("要添加的id是", bookID)
		//获取图书id的信息
		book, err := dao.GetBookByID(bookID)
		if err != nil {
			return
		}
		//获取用户的id
		userID := session.UserID
		//判断当前用户是否有购物车
		fmt.Println("userID", userID)
		cart, _ := dao.GetCartByUserID(userID)
		fmt.Println("cart", cart)
		if cart != nil {
			//当前用户已经有购物车
			//购物车当时是否有这本书
			carItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			fmt.Println("carItem", carItem)
			if carItem != nil {
				//购物车的购物项中已经该图书，只需要将该图书的购物项的count加1即可
				//获取购物车切片中的所有购物项
				item := cart.CartItem
				for _, v := range item {
					//找到当前购物项
					if carItem.Book.ID == v.Book.ID {
						v.Count++
						//更新将数据库中该购物项的图书数量
						dao.UpdateCartItemBookCount(v)
					}
				}
			} else {
				fmt.Println("当前购物车中还木有该图书")
				////购物车的购物项中没有该图书，此时创建一个新的购物项
				//创建购物车中的购物项
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				//将购物项添加到当前cart切片中
				cart.CartItem = append(cart.CartItem, cartItem)
				//将新创建的购物车添加到数据库中
				dao.AddCartItem(cartItem)
			}
			//不管之前购物车中是否有当前图书，都要更新购物车中图书的总数量和总金额
			dao.UpdateCart(cart)
		} else {
			//当前用户还没有购物车
			//cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: session.SessionID,
				UserID: userID,
			}
			//创建购物车中的购物项
			cartItems := []*model.CartItem{}
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: session.SessionID,
			}
			//将购物项添加到切片中
			cartItems = append(cartItems, cartItem)
			cart.CartItem = cartItems
			//将购物车添加到数据库中
			dao.AddCart(cart)
		}
		w.Write([]byte("你刚刚将" + book.Title + "添加到购物车"))
	} else {
		//没有登录
		w.Write([]byte("请先登录"))
	}
}

// 根据用户id获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//判断登录
	_, session := dao.IsLogin(r)
	//获取用id
	userID := session.UserID
	//根据用户id获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	fmt.Println("cart", cart)
	if cart != nil {
		//设置用户名
		//cart.UserName = session.UserName
		session.Cart = cart
		//解析模板文件
		files, err := template.ParseFiles(wd + "/bookstore/view/pages/cart/cart.html")
		if err != nil {
			fmt.Println("files cart.html err", err)
			return
		}
		files.Execute(w, session)
	} else {
		//该用户还没有购物车
		files, err := template.ParseFiles(wd + "/bookstore/view/pages/cart/cart.html")
		if err != nil {
			fmt.Println("files cart.html err", err)
			return
		}
		fmt.Println("session", session)
		files.Execute(w, session)
	}
}

// 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要涮菜的购物车id
	cartID := r.FormValue("cartID")
	fmt.Println(cartID)
	err := dao.DeleteCartByCartID(cartID)
	if err != nil {
		fmt.Println("DeleteCart() err", err)
		return
	}
	//调用GetCartInfo
	GetCartInfo(w, r)
}

// 根据购物项的id删除购物项
func DeleteCartItemByID(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项id
	cartItemID := r.FormValue("cartItemID")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取用户的购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		return
	}
	//获取当前购物车中的购物项
	cartItems := cart.CartItem
	//遍历得到每一个购物项
	for i, item := range cartItems {
		if item.CartItemID == iCartItemID {
			//将当前购物项从切片中移除
			cartItems = append(cartItems[:i], cartItems[i+1:]...)
			//将删除的切片赋给购物项
			cart.CartItem = cartItems
			//删除当前购物项
			err = dao.DeleteCartItemByID(cartItemID)
			if err != nil {
				fmt.Println("DeleteCartItemByID() err", err)
				return
			}
		}
	}
	//更新购物车的信息
	dao.UpdateCart(cart)
	//调用GetCartInfo ，获取当前购物车的信息
	GetCartInfo(w, r)
}

// 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项id
	cartItemID := r.FormValue("cartItemID")
	bookCount := r.FormValue("bookCount")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取用户的购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		return
	}
	//获取当前购物车中的购物项
	cartItems := cart.CartItem
	//遍历得到每一个购物项
	for _, item := range cartItems {
		if item.CartItemID == iCartItemID {
			//将当前购物项中的图书的数量设置为用户输入的值
			item.Count = iBookCount
			//更新数据库中该购物项的count
			dao.UpdateCartItemBookCount(item)

		}
	}
	//更新购物车的信息
	dao.UpdateCart(cart)
	//调用GetCartInfo ，获取当前购物车的信息
	GetCartInfo(w, r)
}
