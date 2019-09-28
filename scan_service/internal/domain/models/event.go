package models

type Event struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}
