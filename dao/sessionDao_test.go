package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestAddSession(t *testing.T) {

	session := &model.Session{
		SessionID: "123123",
		UserName:  "admin",
		UserID:    9,
	}
	err := AddSession(session)
	if err != nil {
		return
	}
}
func TestDeleteSession(t *testing.T) {
	err := DeleteSession("123123")
	if err != nil {
		return
	}
}
func TestGetSession(t *testing.T) {

	session, err := GetSession("715d3aad-6c07-6b1e-6500%!(EXTRA []uint8=[132 132 215 76 205 4])")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(session)
}
