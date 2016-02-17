package email

import (
	"bytes"
	"fmt"
	"text/template"
)

type (
	//Sender basic interface for sending email messages
	Sender interface {
		Send(message *Message) error
	}

	//SenderFactory function for defining an Sender factory
	SenderFactory func() Sender
)

var (
	//NewSender the selected EmailSenderFactory func to be used to instantiate the selected Sender
	NewSender SenderFactory
)

type MessageBody interface {
	String() (string, error)
}

//Message the struct representing an Message
type Message struct {
	From    string
	To      string
	Subject string
	Body    MessageBody
}

//NewMessage a simple factory method for constructing a pointer to an Message
func NewMessage(from string, to string, subject string, body MessageBody) *Message {
	return &Message{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{from:%v, to:%v, subject:%v}", m.From, m.To, m.Subject)
}

type simpleMessageBody struct {
	payload string
}

func NewSimpleMessageBody(body string) MessageBody {
	return simpleMessageBody{
		payload: body,
	}
}

func (body simpleMessageBody) String() (string, error) {
	return body.payload, nil
}

type fileBasedHTMLTemplateMessageBody struct {
	fileName string
	data     interface{}
}

func NewFileBasedHTMLTemplateMessageBody(fileName string, data interface{}) MessageBody {
	return fileBasedHTMLTemplateMessageBody{
		fileName: fileName,
		data:     data,
	}
}

func (f fileBasedHTMLTemplateMessageBody) String() (string, error) {
	t, err := template.ParseFiles(f.fileName)
	if err != nil {
		return "", fmt.Errorf("Failed to parse template from file %s; %v", f.fileName, err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, f.data)
	if err != nil {
		return "", fmt.Errorf("Failed to execute template %s; %v", f.fileName, err)
	}

	return buf.String(), nil
}
