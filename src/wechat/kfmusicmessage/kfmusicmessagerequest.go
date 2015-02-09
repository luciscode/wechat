package kfmusicmessage

type Content struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Musicurl       string `json:"musicurl"`
	Hqmusicurl     string `json:"hqmusicurl"`
	Thumb_media_id string `json:"thumb_media_id"`
}
type Kfmucismessage struct {
	Touser  string  `json:"touser"`
	Msgtype string  `json:"msgtype"`
	Context Content `json:"music"`

	MessageJSON string `json:"-"`
}
