package kfaccount

type Kfaccountresponse struct {
	Errcode     int
	Errmsg      string `json:"errmsg"`
	ReponseJSON string `json:"-"`
}
