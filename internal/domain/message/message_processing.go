package message

type Processing interface {
	Send(msg *Message) error
}
