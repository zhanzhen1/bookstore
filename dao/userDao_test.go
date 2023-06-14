package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("测试登录", testLogin)
	t.Run("add", TestAdduser)
}
func testLogin(t *testing.T) {
	user, err := Login("admin", "123456")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("user", user)
}
func TestAdduser(t *testing.T) {
	err := Adduser("admin", "12345", "admin@qq.com")
	if err != nil {
		fmt.Println("添加失败", err)
		return
	}
}
