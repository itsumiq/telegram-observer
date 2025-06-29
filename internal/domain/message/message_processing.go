package message

type MessageProcessing interface {
	Send(msg *Message) error
}
