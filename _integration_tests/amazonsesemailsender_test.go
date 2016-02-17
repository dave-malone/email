package integrationtests

import (
	"os"
	"testing"

	"github.com/dave-malone/email"
)

func TestAmazonSESSend(t *testing.T) {
	endpoint := os.Getenv("AWS_ENDPOINT")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")

	sender := email.NewAmazonSESSender(endpoint, accessKeyID, secretAccessKey)()

	data := struct {
		Title string
		Items []string
	}{
		Title: "Test page",
		Items: []string{
			"First Item",
			"Second Item",
		},
	}

	body := email.NewFileBasedHTMLTemplateMessageBody("test-html-template.tpl", data)
	err := sender.Send(email.NewMessage(from, to, "Test Email from AmazonSESSender", body))

	if err != nil {
		t.Fatalf("send failed; %v", err)
	}
}
