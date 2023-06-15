package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
)

type User struct {
	Password string `json:"password"`
	Name     string `json:"email"`
}

type SmtpServer struct {
	host string
	port int
}

func main() {
	s := SmtpServer{host: "mail.example.com", port: 465}
	u := User{Name: "yaoshicheng@example.com", Password: "example.com"}
	err := u.verifyEmail(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
func (u *User) verifyEmail(server SmtpServer) error {
	servername := fmt.Sprintf("%s:%d", server.host, server.port)
	host, _, _ := net.SplitHostPort(servername)

	// 设置认证信息。
	auth := smtp.PlainAuth("", u.Name, u.Password, host)

	conn, err := tls.Dial("tcp", servername, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	})
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Quit()

	// Auth
	err = c.Auth(auth)
	if err != nil {
		return err
	}
	return nil
}
