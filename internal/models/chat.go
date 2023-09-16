package models

type Chat struct {
	ID        string
	From      string
	To        string
	Msg       string
	MsgType   string
	Timestamp int64
}

type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Chat Chat   `json:"chat"`
}
