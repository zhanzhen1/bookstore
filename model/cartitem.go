package model

type CartItem struct {
	CartItemID int64   //购物项的id
	Book       *Book   //购物车的图书信息
	Count      int64   //购物车的图书数量
	Amount     float64 //计算得到的
	CartID     string  //购物项属于那个购物车
}

// 获取购物项图书的金额小计
func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
