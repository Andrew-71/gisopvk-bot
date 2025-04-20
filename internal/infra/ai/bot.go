package ai

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Andrew-71/gisopvk-bot/internal/domain"
)

// May this be the one and only time I have to interact with AI of my own volition.
// May the Sun one day shine on an Earth free of it.

type Message struct {
	Role    string `json:"role"` // system/user/assistant,
	Content string `json:"content"`
}

type AiBot struct {
	// messageHistory stores message history by UUIDs.
	// As you can see, it's currently not persistent at all.
	// We are very bad at keeping conversations, okay?
	messageHistory map[string][]Message
	model          string
	apiUrl         string
}

func NewAiBot(apiUrl, model string) *AiBot {
	return &AiBot{
		messageHistory: make(map[string][]Message),
		model:          model,
		apiUrl:         apiUrl,
	}
}

func (b *AiBot) Reply(query domain.Query) (domain.Reply, error) {
	hist, ok := b.messageHistory[query.UUID]
	if !ok {
		hist = InitConversation()
	}

	hist = append(hist, Message{
		Role:    "user",
		Content: query.Body,
	})

	prompt := AiRequest{
		Model:    b.model,
		Messages: hist,
		Stream:   false, // always set to false to reduce complexity
	}
	body, err := json.Marshal(prompt)
	if err != nil {
		return domain.Reply{}, err
	}
	r, err := http.NewRequest("POST", b.apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return domain.Reply{}, err
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil || res.StatusCode != http.StatusOK {
		return domain.Reply{}, err
	}
	defer res.Body.Close()

	var reply AiResponse
	err = json.NewDecoder(res.Body).Decode(&reply)
	if err != nil {
		return domain.Reply{}, err
	}

	hist = append(hist, reply.Message)
	b.messageHistory[query.UUID] = hist

	return domain.Reply{
		UUID: query.UUID,
		Body: reply.Message.Content,
	}, nil
}
