package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 密码的加密和解密
func EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("加密密码失败: %v", err)
	}
	return string(hashPassword), nil
}

func CheckPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
