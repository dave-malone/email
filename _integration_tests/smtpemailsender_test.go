// +build !integration

package integrationtests

import (
	"os"
	"strconv"
	"testing"

	"github.com/dave-malone/email"
)

func TestNewSMTPSender(t *testing.T) {
	server := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	sender := email.NewSMTPSender(server, port, username, password)

	if sender == nil {
		t.Fatal("NewSMTPSender returned a nil value")
	}
}

func TestSMTPSend(t *testing.T) {
	server := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")

	sender := email.NewSMTPSender(server, port, username, password)()

	err := sender.Send(email.NewMessage(from, to, "Test Email from SMTPSender", "Body Text", "Body HTML"))

	if err != nil {
		t.Fatalf("send failed; %v", err)
	}
}
