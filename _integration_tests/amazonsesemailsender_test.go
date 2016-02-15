package integrationtests

import (
	"os"
	"testing"

	"github.com/dave-malone/email"
)

func TestNewAmazonSESSender(t *testing.T) {
	endpoint := os.Getenv("AWS_ENDPOINT")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	sender := email.NewAmazonSESSender(endpoint, accessKeyID, secretAccessKey)

	if sender == nil {
		t.Fatal("NewAmazonSESSender returned a nil value")
	}
}

func TestAmazonSESSend(t *testing.T) {
	endpoint := os.Getenv("AWS_ENDPOINT")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")

	sender := email.NewAmazonSESSender(endpoint, accessKeyID, secretAccessKey)()

	err := sender.Send(email.NewMessage(from, to, "Test Email from AmazonSESSender", "Body Text", "Body HTML"))

	if err != nil {
		t.Fatalf("send failed; %v", err)
	}
}
