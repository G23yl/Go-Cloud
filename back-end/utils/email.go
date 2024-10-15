package utils

import (
	"bytes"
	"crypto/tls"
	"disk/config"
	"disk/database"
	"fmt"
	"html/template"
	"time"

	"github.com/thanhpk/randstr"
	"gopkg.in/gomail.v2"
)

// 生成验证码
func GenerateRandomCode(num int) string {
	return randstr.Dec(num)
}

// 发送验证邮件
func SendEmail(email string) error {
	randomCode := GenerateRandomCode(4)
	codeVal := struct {
		Code string
	}{randomCode}
	randomCodeExpired := 5 * time.Minute
	// 解析邮件模板
	m := gomail.NewMessage()
	m.SetHeader("From", config.GetEmailFrom())
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码")
	template, err := template.ParseFiles("utils/static/email.html")
	if err != nil {
		fmt.Println("邮件模板parse失败")
		return err
	}
	var body bytes.Buffer
	if err := template.Execute(&body, codeVal); err != nil {
		fmt.Println("邮件模板execute失败")
		return err
	}
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(config.GetEmailHost(), config.GetEmailPort(), config.GetEmailFrom(), config.GetEmailPassword())
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// 将验证码存入redis
	if err := database.Set(email, randomCode, randomCodeExpired); err != nil {
		fmt.Println("验证码存储失败")
		return err
	}
	return d.DialAndSend(m)
}
