package ai

import "time"

type AiRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type AiResponse struct {
	Model   string    `json:"model"`
	Created time.Time `json:"created_at"`
	Message Message   `json:"message"`
	Done    bool      `json:"done"`
	/*
		We could also support these:
		"total_duration": 5191566416,
		"load_duration": 2154458,
		"prompt_eval_count": 26,
		"prompt_eval_duration": 383809000,
		"eval_count": 298,
		"eval_duration": 4799921000
	*/
}
