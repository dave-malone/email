package email

import "testing"

func TestNewNoopSender(t *testing.T) {
	sender := NewNoopSender()

	if sender == nil {
		t.Fatal("newNoopSender returned a nil value")
	}
}

func TestNoopSend(t *testing.T) {
	NewSender = NewNoopSender
	sender := NewSender()

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

	body := NewSimpleMessageBody("Test simple message body")
	err := sender.Send(NewMessage("from@test.com", "to@test.com", "Subject", body))

	if err != nil {
		t.Fatal("noopSender.send should always return a nil error")
	}
}
