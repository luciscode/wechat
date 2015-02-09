package kfvoicemessage

type Content struct {
	Content string `json:"media_id"`
}
type Kfvoicemessage struct {
	Touser  string  `json:"touser"`
	Msgtype string  `json:"msgtype"`
	Context Content `json:"voice"`

	MessageJSON string `json:"-"`
}
