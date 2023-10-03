package embedding

type Request struct {
	InputText string `json:"inputText"`
}

type Response struct {
	Embedding           []float32 `json:"embedding"`
	InputTextTokenCount int       `json:"inputTextTokenCount"`
}
