package jswechat

import (
	"wechat/wsign"
)

type Sha1message struct {
	Jsapi_ticket string `json:"jsapi_ticket"`
	NonceStr     string `json:"noncestr"`
	Timestamp    string `json:"timestamp"`
	Url          string `json:"url"`
}
type Config struct {
	AppId     string `json:"appId"`
	Debug     bool   `json:"debug"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	JsApiList string `json:"timestamp"`
}

func (v Sha1message) SignSHA1() string {
	sign, _ := wsign.SignSHA1(&v)
	return sign
}
