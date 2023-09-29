package stabilityai

import "encoding/base64"

type Request struct {
	TextPrompts []TextPrompt `json:"text_prompts"`
	CfgScale    float64      `json:"cfg_scale"`
	Steps       int          `json:"steps"`
	Seed        int          `json:"seed"`
}

type TextPrompt struct {
	Text string `json:"text"`
}

type Response struct {
	Result    string     `json:"result"`
	Artifacts []Artifact `json:"artifacts"`
}

type Artifact struct {
	Base64       string `json:"base64"`
	FinishReason string `json:"finishReason"`
}

func (a *Artifact) DecodeImage() ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(a.Base64)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
