package cohere

type Request struct {
	Prompt            string           `json:"prompt"`
	Temperature       float64          `json:"temperature,omitempty"`
	P                 float64          `json:"p,omitempty"`
	K                 float64          `json:"k,omitempty"`
	MaxTokens         int              `json:"max_tokens,omitempty"`
	StopSequences     []string         `json:"stop_sequences,omitempty"`
	ReturnLikelihoods ReturnLikelihood `json:"return_likelihoods,omitempty"`
	Stream            bool             `json:"stream,omitempty"`
	NumGenerations    int              `json:"num_generations,omitempty"`
}

type ReturnLikelihood string

const (
	Generation ReturnLikelihood = "GENERATION"
	All        ReturnLikelihood = "ALL"
	None       ReturnLikelihood = "NONE"
)
