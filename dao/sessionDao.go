package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// addSession向数据库中添加session
func AddSession(session *model.Session) error {
	sqlStr := "insert into session values (?,?,?)"
	_, err := utils.DB.Exec(sqlStr, session.SessionID, session.UserName, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

// 删除数据库session
func DeleteSession(sessionID string) error {
	sqlStr := "delete from session where session_id = ?"
	_, err := utils.DB.Exec(sqlStr, sessionID)
	if err != nil {
		fmt.Println("DeleteSession() err", err)
		return err
	}
	return nil
}

// 获取seesion
func GetSession(sessionID string) (*model.Session, error) {
	sqlStr := "select * from session where session_id = ?"
	prepare, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := prepare.QueryRow(sessionID)
	session := &model.Session{}
	err = row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	if err != nil {
		fmt.Println("session scan err", err)
		return nil, err
	}
	//row := utils.DB.QueryRow(sqlStr, sessionID)
	//session := &model.Session{}
	//err := row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	//if err != nil {
	//	fmt.Println("session scan err", err)
	//	return nil, err
	//}
	return session, nil
}

// 判断是否已经登录
func IsLogin(ctx *gin.Context) (bool, *model.Session) {
	cookie, _ := ctx.Cookie("user")
	fmt.Println("cookie :", cookie)
	if cookie != "" {
		cookieValue := cookie
		fmt.Println("cookieValue:", cookieValue)
		session, _ := GetSession(cookieValue)
		fmt.Println("session:", session)
		if session != nil {
			//已经登录了
			return true, session
		}
	}
	return false, nil
}
