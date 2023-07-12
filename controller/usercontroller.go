package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 优化登录
func Login() (handleFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		//判断是否已经登录
		faly, _ := dao.IsLogin(context)
		if faly {
			//已经登录
			GetPageBookByPrice()
		} else {
			//获取用户名密码
			username := context.PostForm("username")
			password := context.PostForm("password")
			//调用验证用户名和密码的方法
			user, _ := dao.Login(username, password)
			//if err != nil {
			//	fmt.Println("Login() err ", err)
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
				//将cookie发送给浏览器
				_, err = context.Cookie("user")
				if err != nil {
					context.SetCookie("user", uuid, 3600, "/", "localhost", false, true)
				}
				fmt.Println("cookie：", uuid)
				context.HTML(http.StatusOK, "login_success.html", user)
			} else {
				context.HTML(http.StatusOK, "login.html", "用户名和密码不正确")
				fmt.Println("用户名和密码不正确")
			}
		}
	}

}

// gin 优化注销
func Logout() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		//获取cookie
		cookie, err := context.Cookie("user")
		if err != nil {
			return
		}
		fmt.Println("cookie", cookie)
		if cookie != "" {
			//删除数据库对应得session中的用户名
			err := dao.DeleteSession(cookie)
			if err != nil {
				fmt.Println("DeleteSession() err", err)
				return
			}
			//设置cookie失效  maxage=0 未设置 maxage<0 失效 maxage>0 存在
			context.SetCookie("user", cookie, -1, "/", "localhost", false, true)
			//将修改之后得cookie发送给浏览器
		}
		//去首页
		context.HTML(http.StatusOK, "index.html", nil)
	}
}

// gin 优化注册

func Register() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		//获取用户名密码
		username := context.PostForm("username")
		password := context.PostForm("password")
		repwd := context.PostForm("repwd")
		email := context.PostForm("email")
		//调用验证用户名和密码的方法
		user, err := dao.Register(username)
		if err != nil {
			fmt.Println("Register() err", err)
			return
		}
		fmt.Println("获取的user是", user)
		if user == nil {
			//用户名存在
			context.HTML(http.StatusOK, "regist.html", "用户名已存在")
		} else {
			if repwd != password {
				context.HTML(http.StatusOK, "regist.html", "密码不一致")
				return
			}
			context.HTML(http.StatusOK, "regist_success.html", nil)
			//将用户保存在数据库中
			err = dao.Adduser(username, password, email)
			if err != nil {
				fmt.Println("注册err", err)
				return
			}
			fmt.Println("注册成功")
		}
	}

}

// 获取username 判断用户名是否存在
func CheckUserName() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		//获取用户名密码
		username := context.PostForm("username")
		//调用验证用户名和密码的方法
		user, _ := dao.Register(username)
		fmt.Println("获取的user是", user)
		if user != nil {
			//用户名存在
			context.Writer.Write([]byte("用户名存在"))
		} else {
			//用户名可用
			context.Writer.Write([]byte("用户名可用"))

		}
	}

}

//原生go写的项目
// 登录
//func Login1(w http.ResponseWriter, r *http.Request) {
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	//判断是否已经登录
//	faly, _ := dao.IsLogin(r)
//	if faly {
//		//已经登录
//		GetPageBookByPrice()
//	} else {
//		//获取用户名密码
//		username := r.PostFormValue("username")
//		password := r.PostFormValue("password")
//		//调用验证用户名和密码的方法
//		user, _ := dao.Login(username, password)
//		//if err != nil {
//		//	fmt.Println("denglu err ", err)
//		//	return
//		//}
//		fmt.Println("获取的user是", user)
//		if user != nil {
//			//生成uuid
//			uuid := utils.CreateUUID()
//			//创建一个session
//			session := &model.Session{
//				SessionID: uuid,
//				UserName:  username,
//				UserID:    user.ID,
//			}
//			//将创建得seesion保存到数据库中
//			err := dao.AddSession(session)
//			if err != nil {
//				return
//			}
//			//创建cookie
//			cookie := &http.Cookie{
//				Name:     "user",
//				Value:    uuid,
//				HttpOnly: true,
//			}
//			//将cookie发送给浏览器
//			http.SetCookie(w, cookie)
//			files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/login_success.html")
//			if err != nil {
//				fmt.Println("跳转err", err)
//				return
//			}
//			files.Execute(w, user)
//
//		} else {
//			files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/login.html")
//			if err != nil {
//				return
//			}
//			files.Execute(w, "用户名和密码不正确")
//			fmt.Println("用户名和密码不正确")
//		}
//	}
//}

// 注册用户

//	func Register(w http.ResponseWriter, r *http.Request) {
//		wd, err := os.Getwd()
//		if err != nil {
//			log.Fatal(err)
//		}
//		//获取用户名密码
//		username := r.PostFormValue("username")
//		password := r.PostFormValue("password")
//		repwd := r.PostFormValue("repwd")
//		email := r.PostFormValue("email")
//		//调用验证用户名和密码的方法
//		user, _ := dao.Register(username)
//		fmt.Println("获取的user是", user)
//
//		if user != nil {
//			//用户名存在
//			files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/regist.html")
//			if err != nil {
//				fmt.Println("跳转err", err)
//				return
//			}
//			files.Execute(w, "用户名已存在")
//		} else {
//			if repwd != password {
//				files, _ := template.ParseFiles(wd + "/bookstore/view/pages/user/regist.html")
//				files.Execute(w, "密码不一致")
//				return
//			}
//			files, err := template.ParseFiles(wd + "/bookstore/view/pages/user/regist_success.html")
//			if err != nil {
//				fmt.Println("跳转err", err)
//				return
//			}
//			//将用户保存在数据库中
//			err = dao.Adduser(username, password, email)
//			if err != nil {
//				fmt.Println("注册err", err)
//				return
//			}
//			files.Execute(w, "")
//			fmt.Println("注册成功")
//		}
//	}

// 注销
//func Logout1(w http.ResponseWriter, r *http.Request) {
//	//获取cookie
//	cookie, err := r.Cookie("user")
//	if err != nil {
//		return
//	}
//	fmt.Println("cookie", cookie)
//	if cookie != nil {
//		cookieValue := cookie.Value
//		//删除数据库对应得session中的用户名
//		err := dao.DeleteSession(cookieValue)
//		if err != nil {
//			return
//		}
//		//设置cookie失效  maxage=0 未设置 maxage<0 失效 maxage>0 存在
//		cookie.MaxAge = -1
//		//将修改之后得cookie发送给浏览器
//		http.SetCookie(w, cookie)
//		//去首页
//		GetPageBookByPrice()
//	}
//}
