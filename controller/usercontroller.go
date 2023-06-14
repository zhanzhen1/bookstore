package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// 登录
func Login(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//判断是否已经登录
	faly, _ := dao.IsLogin(r)
	if faly {
		//已经登录
		GetPageBookByPrice(w, r)
	} else {
		//获取用户名密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用验证用户名和密码的方法
		user, _ := dao.Login(username, password)
		//if err != nil {
		//	fmt.Println("denglu err ", err)
		//	return
		//}
		fmt.Println("获取的user是", user)
		if user != nil {
			//生成uuid
			uuid := utils.CreateUUID()
			//创建一个session
			session := &model.Session{
				SessionID: uuid,
				UserName:  username,
				UserID:    user.ID,
			}
			//将创建得seesion保存到数据库中
			err := dao.AddSession(session)
			if err != nil {
				return
			}
			//创建cookie
			cookie := &http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, cookie)
			files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/login_success.html")
			if err != nil {
				fmt.Println("跳转err", err)
				return
			}
			files.Execute(w, user)

		} else {
			files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/login.html")
			if err != nil {
				return
			}
			files.Execute(w, "用户名和密码不正确")
			fmt.Println("用户名和密码不正确")
		}
	}
}

// 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		return
	}
	fmt.Println("cookie", cookie)
	if cookie != nil {
		cookieValue := cookie.Value
		//删除数据库对应得session中的用户名
		err := dao.DeleteSession(cookieValue)
		if err != nil {
			return
		}
		//设置cookie失效  maxage=0 未设置 maxage<0 失效 maxage>0 存在
		cookie.MaxAge = -1
		//将修改之后得cookie发送给浏览器
		http.SetCookie(w, cookie)
		//去首页
		GetPageBookByPrice(w, r)
	}
}

// 注册用户
func Register(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//获取用户名密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	repwd := r.PostFormValue("repwd")
	email := r.PostFormValue("email")
	//调用验证用户名和密码的方法
	user, _ := dao.Register(username)
	fmt.Println("获取的user是", user)

	if user != nil {
		//用户名存在
		files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/regist.html")
		if err != nil {
			fmt.Println("跳转err", err)
			return
		}
		files.Execute(w, "用户名已存在")
	} else {
		if repwd != password {
			files, _ := template.ParseFiles(wd + "/bookstore/view/pages/user/regist.html")
			files.Execute(w, "密码不一致")
			return
		}
		files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/regist_success.html")
		if err != nil {
			fmt.Println("跳转err", err)
			return
		}
		//将用户保存在数据库中
		err = dao.Adduser(username, password, email)
		if err != nil {
			fmt.Println("注册err", err)
			return
		}
		files.Execute(w, "")
		fmt.Println("注册成功")
	}
}

// 获取username 判断用户名是否存在
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户名密码
	username := r.PostFormValue("username")
	//调用验证用户名和密码的方法
	user, _ := dao.Register(username)
	fmt.Println("获取的user是", user)
	if user != nil {
		//用户名存在
		w.Write([]byte("用户名存在"))
	} else {
		//用户名可用
		w.Write([]byte("用户名可用"))

	}
}