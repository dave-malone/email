// +build !integration

package integrationtests

import (
	"os"
	"strconv"
	"testing"

	"github.com/dave-malone/email"
)

func TestSMTPSend(t *testing.T) {
	server := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")

	sender := email.NewSMTPSender(server, port, username, password)()

	body := email.NewSimpleMessageBody("Simple Message Body")
	err := sender.Send(email.NewMessage(from, to, "Test Email from SMTPSender", body))

	if err != nil {
		t.Fatalf("send failed; %v", err)
	}
}
