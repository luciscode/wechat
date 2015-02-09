package kftextmessages

type Content struct {
	Content string `json:"content"`
}
type Kftextmessage struct {
	Touser  string  `json:"touser"`
	Msgtype string  `json:"msgtype"`
	Context Content `json:"text"`

	MessageJSON string `json:"-"`
}
