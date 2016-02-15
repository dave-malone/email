package email

import "fmt"

type (
	//Sender basic interface for sending email messages
	Sender interface {
		Send(message *Message) error
	}

	//SenderFactory function for defining an Sender factory
	SenderFactory func() Sender
)

var (
	//NewSenderFactory the selected EmailSenderFactory func to be used
	NewSenderFactory SenderFactory
)

//Message the struct representing an Message
type Message struct {
	From, To, Subject, BodyText, BodyHTML string
}

//NewMessage a simple factory method for constructing a pointer to an Message
func NewMessage(from, to, subject, bodyText, bodyHTML string) *Message {
	return &Message{
		From:     from,
		To:       to,
		Subject:  subject,
		BodyText: bodyText,
		BodyHTML: bodyHTML,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{from:%v, to:%v, subject:%v}", m.From, m.To, m.Subject)
}
