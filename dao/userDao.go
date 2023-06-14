package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

// 登录
func Login(username string, password string) (*model.User, error) {
	sqlStr := "select * from user where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 注册用户
func Register(username string) (*model.User, error) {
	sqlStr := "select * from user where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 添加用户
func Adduser(username string, password string, email string) error {
	sqlStr := "INSERT INTO user (username,password,eamil) VALUES(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("add err", err)
		return err
	}
	return nil
}
