package card

type Cardreponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Card_id string `json:"card_id"`

	ReponseXML string `json:"-"`
}
