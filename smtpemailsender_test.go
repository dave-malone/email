package email

import "testing"

func TestNewMTPSender(t *testing.T) {
	sender := NewSMTPSender("", -1, "", "")

	if sender == nil {
		t.Fatal("NewSMTPSender returned a nil value")
	}
}

func TestSMTPSend(t *testing.T) {
	NewSenderFactory = NewSMTPSender("", -1, "", "")
	sender := NewSenderFactory()

	err := sender.send(NewMessage("from@test.com", "to@test.com", "Subject", "Body Text", "Body HTML"))

	if err != nil {
		t.Fatalf("send failed; %v", err)
	}
}
