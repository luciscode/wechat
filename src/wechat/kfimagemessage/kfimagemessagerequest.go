package kfimagemessage

type Content struct {
	Media_id string `json:"media_id"`
}
type Kfimagemessage struct {
	Touser  string  `json:"touser"`
	Msgtype string  `json:"msgtype"`
	Context Content `json:"image"`

	MessageJSON string `json:"-"`
}
