package email

import "github.com/xchapter7x/lo"

type noopSender struct {
}

//NewNoopSender factory method for a Sender implementation that simply logs out email messages
func NewNoopSender() Sender {
	return noopSender{}
}

func (sender noopSender) Send(message *Message) error {
	lo.G.Debugf("message received: %v\n", message)

	return nil
}
