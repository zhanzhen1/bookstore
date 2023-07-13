package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
)

var user *model.User
var userList []*model.User

// 登录
func Login(username string, password string) (*model.User, error) {
	if err := utils.Db1.Where("username = ? and password = ?", username, password).
		Find(&user).Error; err != nil {
		log.Fatal("login() err", err)
	}
	return user, nil
}

// 注册用户
func Register(username string) (*model.User, error) {
	if err := utils.Db1.Where("username = ?", username).Find(&user).Error; err != nil {
		log.Fatal("Register() err", err)
	}
	return user, nil
}

// 添加用户
func Adduser(username string, password string, email string) error {
	user = &model.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	if err := utils.Db1.Create(&user).Error; err != nil {
		log.Fatal("Create() err", err)
	}
	return nil
}
