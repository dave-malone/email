package email

import "testing"

func TestNewAmazonSESSender(t *testing.T) {
	sender := NewAmazonSESSender("", "", "")

	if sender == nil {
		t.Fatal("NewAmazonSESSender returned a nil value")
	}
}

func TestAmazonSESSend(t *testing.T) {
	NewSenderFactory = NewAmazonSESSender("", "", "")
	sender := NewSenderFactory()

	err := sender.send(NewMessage("from@test.com", "to@test.com", "Subject", "Body Text", "Body HTML"))

	if err != nil {
		t.Fatalf("send failed; %v", err)
	}
}
