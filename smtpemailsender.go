package email

import (
	"fmt"
	"net/smtp"
)

type smtpSender struct {
	host     string
	port     int
	username string
	password string
}

//NewSMTPSender returns an implementation of email.Sender for sending email messages over smtp
func NewSMTPSender(host string, port int, username string, password string) SenderFactory {
	return func() Sender {
		return smtpSender{
			host:     host,
			port:     port,
			username: username,
			password: password,
		}
	}
}

func (sender smtpSender) Send(message *Message) error {
	auth := smtp.PlainAuth("", sender.username, sender.password, sender.host)

	addr := fmt.Sprintf("%s:%d", sender.host, sender.port)
	to := []string{message.To}

	body, err := message.Body.String()

	if err == nil {
		msg := []byte(fmt.Sprintf("To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", message.To, message.Subject, body))
		err = smtp.SendMail(addr, auth, message.From, to, msg)
	}

	return err
}
