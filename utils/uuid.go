package utils

import (
	"crypto/rand"
	"fmt"
)

// 获取uuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		return ""
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0x7F) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
