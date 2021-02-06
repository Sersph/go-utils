package email

import (
	"fmt"
	"net/smtp"
	"testing"
)

func TestIdUtils(t *testing.T) {
	auth := smtp.PlainAuth("", "18934086807@163.com", "ENKCEDSUODGKDAGF", "smtp.163.com")
	msg := []byte("To: 2675835744@qq.com\r\nFrom: 18934086807@163.com>\r\nSubject: " + "\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n" + "Hello World!")

	err := smtp.SendMail("smtp.163.com:25", auth, "18934086807@163.com", []string{"2675835744@qq.com"}, msg)
	fmt.Println(err)
}
