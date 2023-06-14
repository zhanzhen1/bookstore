package model

type Cart struct {
	CartID      string
	CartItem    []*CartItem //购物车所有的购物项
	TotalCount  int64       //购物车图书的总量，计算得到
	TotalAmount float64     //购物车的总金额
	UserID      int         //当前购物车所属的用户
	UserName    string
}

// 获取购物车中图书的总数
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中购物项切片
	for _, item := range cart.CartItem {
		totalCount = totalCount + item.Count
	}
	return totalCount
}

// 获取购物车中图书的金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历购物车中购物项切片
	for _, item := range cart.CartItem {
		totalAmount = totalAmount + item.GetAmount()
	}
	return totalAmount
}
