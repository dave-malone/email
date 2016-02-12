package email

import "errors"

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

func (sender smtpSender) send(message *Message) error {
	return errors.New("Not Yet Implemented")
}
