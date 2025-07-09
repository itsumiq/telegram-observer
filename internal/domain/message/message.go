package message

import "strings"

type Message struct {
	ChatID   int64
	Text     string
	Buttons  []Button
	FilePath string
}

type Button struct {
	Text string
	Data string
}

func New(chatID int64, text string, filePath string) *Message {
	text = convertToMarkdown(text)
	return &Message{ChatID: chatID, Text: text, FilePath: filePath}
}

func (m *Message) AddButton(text string, data string) {
	m.Buttons = append(m.Buttons, Button{Text: text, Data: data})
}

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
