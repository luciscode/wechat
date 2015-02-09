package card

type Updateshakecardresponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`

	ReponseXML string `json:"-"`
}
