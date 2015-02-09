package kftextmessages

import (
	"wechat/wjson"
)

func Getkftextmessage(touser, content string) Kftextmessage {
	v := Kftextmessage{}
	v.Msgtype = "text"
	v.Touser = touser

	c := Content{}
	c.Content = content
	v.Context = c
	v.MessageJSON, _ = wjson.Encodestruct(v)
	return v
}
