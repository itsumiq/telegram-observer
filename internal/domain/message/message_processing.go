package message

// Processing represents contract that message processings should implement.
type Processing interface {
	// Send sends text message and returns error.
	Send(msg *Message) error

	// SendPhotoMessage sends photo message and returns error.
	SendPhotoMessage(msg *Message) error
}
