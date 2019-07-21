package constants

type GenericMessage struct {
	Result interface{} `json:"result"`
	Status int         `json:"status"`
}

const (
	SuccessStatus = 1
	ErrorStatus   = 0
)
