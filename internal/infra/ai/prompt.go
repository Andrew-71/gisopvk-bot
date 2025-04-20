package ai

import _ "embed"

//go:embed prompt.txt
var prompt string

func InitConversation() []Message {
	hist := make([]Message, 0)
	hist = append(hist, Message{
		Role:    "system",
		Content: prompt,
	})
	hist = append(hist, Message{
		Role:    "assistant",
		Content: "Привет! Как я могу вам помочь?",
	})
	return hist
}
