package kfvideomessage

type Content struct {
	Media_id    string `json:"media_id"`
	Tedia_id    string `json:"thumb_media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Kfvideomessage struct {
	Touser  string  `json:"touser"`
	Msgtype string  `json:"msgtype"`
	Context Content `json:"video"`

	MessageJSON string `json:"-"`
}
