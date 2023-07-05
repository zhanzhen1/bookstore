package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ReturnErr 返回错误页面
func ReturnErr(ctx *gin.Context, err error) {
	ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"error": err.Error(),
	})
}

// GetPageNo 从请求中获取当前所在页
func GetPageNo(ctx *gin.Context) int64 {
	//转换成int64
	i, _ := strconv.ParseInt(ctx.Query("pageNo"), 10, 64)
	if i <= 0 {
		return 1
	}
	return i
}
