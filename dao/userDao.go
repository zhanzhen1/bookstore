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
	utils.Db1.Where("username = ? and password = ?", username, password).Find(&user)
	return user, nil
}

// 注册用户
func Register(username string) (*model.User, error) {
	utils.Db1.Where("username = ?", username).Find(&user)
	return user, nil
}

// 添加用户
func Adduser(username string, password string, email string) error {
	//sqlStr := "INSERT INTO user (username,password,eamil) VALUES(?,?,?)"
	//_, err := utils.DB.Exec(sqlStr, username, password, email)
	//if err != nil {
	//	fmt.Println("add err", err)
	//	return err
	//}
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
