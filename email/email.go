package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/lincanjian/golink"
)

func Create(link string) (email *golink_email, err error) {

	arr, err := golink.Link_Verify("email", 5, link)
	if err != nil {
		return
	}

	email = &golink_email{
		smtpServer:  arr[1],
		smtpPort:    arr[2],
		from:        arr[3],
		licenseCode: arr[4],
	}

	accessAddress := email.smtpServer + ":" + email.smtpPort
	// 连接SMTP服务器
	auth := smtp.PlainAuth("", email.from, email.licenseCode, email.smtpServer)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         email.smtpServer,
	}

	email.emailClient, err = smtp.Dial(accessAddress)
	if err != nil {
		return
	}

	// 开启TLS加密
	if err = email.emailClient.StartTLS(tlsConfig); err != nil {
		return
	}

	// 登录SMTP服务器
	if err = email.emailClient.Auth(auth); err != nil {
		return
	}
	return
}

func (l *golink_email) SendEmail(to, subject, body string) error {
	err := l.SendMultipleEmail([]string{to}, subject, body)
	return err
}

func (l *golink_email) SendMultipleEmail(to []string, subject, body string) error {
	// 邮件头部信息
	headers := make(map[string]string)
	headers["From"] = l.from
	headers["Subject"] = subject
	if len(to) == 1 {
		headers["To"] = to[0]
	}

	if len(to) > 1 {
		headers["To"] = strings.Join(to, ",")
	}

	// 将邮件头部信息转换为RFC 822格式
	var message string
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 发送邮件
	if err := l.emailClient.Mail(l.from); err != nil {
		return err
	}

	for _, recipient := range to {
		if err := l.emailClient.Rcpt(recipient); err != nil {
			fmt.Println("添加收件人失败：", recipient, err)
		}
	}

	// 发送邮件
	w, err := l.emailClient.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	return err
}
