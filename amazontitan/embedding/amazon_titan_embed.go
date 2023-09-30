package embedding

type Request struct {
	InputText string `json:"inputText"`
}

type Response struct {
	Embedding           []float64 `json:"embedding"`
	InputTextTokenCount int       `json:"inputTextTokenCount"`
}
