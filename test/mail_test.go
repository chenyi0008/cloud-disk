package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Chen <chenyikeji0000@163.com>"
	e.To = []string{"71563698@qq.com"}
	e.Subject = "验证码发送测试："
	e.HTML = []byte("你的验证码为：<h1>123123</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", define.MailAccount, define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
