package model

// 分页
type Page struct {
	Book        []*Book
	PageNo      int64 //当前页
	PageSize    int64 //每页显示几条
	TotalPageNo int64 //总页数
	TotalRecord int64
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

// 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// 获取上一页数据
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

// 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
