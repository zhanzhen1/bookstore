package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetManager() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, "manager.html", nil)
	}
}

// 跳转登录
func GetLogin() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	}
}

// 跳转注册
func GetRegister() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, "regist.html", nil)
	}
}

// 跳转购物车
func GetCart() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, "cart.html", nil)
	}
}
