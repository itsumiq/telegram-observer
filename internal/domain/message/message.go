package message

import "strings"

// Message represents  message domain entity.
type Message struct {
	ChatID   int64
	Text     string
	Buttons  []Button
	FilePath string
}

// Button represent Message field.
type Button struct {
	Text string
	Data string
}

// New creates new message instance.
func New(chatID int64, text string, filePath string) *Message {
	text = convertToMarkdown(text)
	return &Message{ChatID: chatID, Text: text, FilePath: filePath}
}

// AddButton adds button to message.
func (m *Message) AddButton(text string, data string) {
	m.Buttons = append(m.Buttons, Button{Text: text, Data: data})
}

// convertToMarkdown converts text markup to markdown and returns it.
func convertToMarkdown(text string) string {
	specChars := []string{
		"_",
		"[",
		"]",
		"(",
		")",
		"~",
		"`",
		">",
		"#",
		"+",
		"-",
		"=",
		"|",
		"{",
		"}",
		".",
		"!",
	}

	for _, char := range specChars {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}

	return text
}
