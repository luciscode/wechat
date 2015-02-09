package kfvoicemessage

import (
	"wechat/wjson"
)

func Getkfvoicemessage(touser, media_id string) Kfvoicemessage {
	v := Kfvoicemessage{}
	v.Msgtype = "voice"
	v.Touser = touser

	c := Content{}
	c.Content = media_id
	v.Context = c
	v.MessageJSON, _ = wjson.Encodestruct(v)
	return v
}
