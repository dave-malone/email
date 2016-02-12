package email

import "testing"

func TestNewNoopSender(t *testing.T) {
	sender := NewNoopSender()

	if sender == nil {
		t.Fatal("newNoopSender returned a nil value")
	}
}

func TestNoopSend(t *testing.T) {
	NewSenderFactory = NewNoopSender
	sender := NewSenderFactory()

	err := sender.send(NewMessage("from@test.com", "to@test.com", "Subject", "Body Text", "Body HTML"))

	if err != nil {
		t.Fatal("noopSender.send should always return a nil error")
	}
}
