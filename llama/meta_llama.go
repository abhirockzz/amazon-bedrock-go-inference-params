package llama

import "strings"

type Request struct {
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	MaxGenLen   int     `json:"max_gen_len,omitempty"`
}

type Response struct {
	Generation           string `json:"generation"`
	PromptTokenCount     int    `json:"prompt_token_count"`
	GenerationTokenCount int    `json:"generation_token_count"`
	StopReason           string `json:"stop_reason"`
}

func (r Response) GetResponseString() string {
	return strings.TrimPrefix(r.Generation, "\n\n")
}
