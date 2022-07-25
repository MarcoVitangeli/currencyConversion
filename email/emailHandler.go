package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"path/filepath"
	"text/template"
)

type EmailHandler struct {
	From     string
	To       []string
	Password string
	Host     string
	Port     string
}

func getPath() (string, error) {
	return filepath.Abs("./email/ASSET/template.html")
}

func (handler *EmailHandler) SendCurrInfo(data float32) {
	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		println(err)
	}
	c, err := smtp.NewClient(conn, handler.Host)
	if err != nil {
		println(err)
	}

	tlsconfig := &tls.Config{
		ServerName: handler.Port,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
	}

	auth := LoginAuth(handler.From, handler.Password)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: USD Daily conversion \n%s\n\n", mimeHeaders)))

	path, err := getPath()

	if err != nil {
		log.Panic(err)
	}

	t, err := template.ParseFiles(path)

	if err != nil {
		log.Panic(err)
	}

	err = t.Execute(&body, struct {
		Info float32
	}{
		Info: data,
	})

	if err != nil {
		println(err)
	}

	err = smtp.SendMail(handler.Host+":"+handler.Port, auth, handler.From, handler.To, body.Bytes())
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
