package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

//define TestSendEmail function
func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "zhangke_2021@126.com"
	e.To = []string{"zhangxiaocong69@hotmail.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//tsl port 587  common port 25
	err := e.SendWithTLS("smtp.126.com:587", smtp.PlainAuth("", "zhangke_2021@126.com", "YWKWXQBYMWCDJIHV", "smtp.126.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.126.com"})
	if err != nil {
		panic(err)
	}
}
