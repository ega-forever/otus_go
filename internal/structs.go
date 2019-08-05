package internal

type Product struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type GenericMessage struct {
	Result interface{} `json:"result"`
	Status int         `json:"status"`
}

const (
	SuccessStatus = 1
	ErrorStatus   = 0
)
