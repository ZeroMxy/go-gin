package email

import (
	"crypto/tls"
	"fmt"
	"go-gin/config"
	"go-gin/core/log"
	"net/smtp"
)

type Email struct {}

type Message struct {
	Subject string
	ContentType string
	Content string
}

var (
	host = config.App["smtpHost"]
	port = config.App["smtpPort"]
	username = config.App["smtpUsername"]
	password = config.App["smtpPassword"]
)

// 发送邮件
// recipients 收件人
// subject 主题
// message 消息
func Send (recipients []string, subject string, message Message) bool {

	conn := conn()
	defer conn.Quit()
	
	// 发起 STARTTLS 协议
	if ok, _ := conn.Extension("STARTTLS"); ok {
		if err := conn.StartTLS(&tls.Config {
			ServerName: host,
		}); err != nil {
			log.Errorf("Email startTLS protocol error: %v", err)
			return false
		}
	}

	// 认证
	if err := conn.Auth(auth()); err != nil {
		log.Errorf("Email auth error: %v", err)
		return false
	}

	// 发送
	from := username
	for _, recipient := range recipients {
		content := fmt.Sprintf(
			"To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%v",
			recipient, from, message.Subject, message.ContentType, message.Content,
		)
		log.Debug(content)
		if err := smtp.SendMail(host + ":" + port, auth(), from, []string{recipient}, []byte(content)); err != nil {
			log.Error("Email send error: %v", err)
		}
	}
	

	return true
}

// SMTP 认证
func auth () smtp.Auth {
	return smtp.PlainAuth("", username, password, host)
}

// 连接 SMTP 服务器
func conn () *smtp.Client {

	conn, err := smtp.Dial(host + ":" + port)
	if err != nil {
		log.Errorf("Email conn error: ", err)
		return nil
	}

	return conn
}
