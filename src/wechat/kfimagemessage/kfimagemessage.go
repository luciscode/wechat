package kfimagemessage

import (
	"wechat/wjson"
)

func Getkfimagemessage(touser, media_id string) Kfimagemessage {
	v := Kfimagemessage{}
	v.Msgtype = "image"
	v.Touser = touser

	c := Content{}
	c.Media_id = media_id
	v.Context = c
	v.MessageJSON, _ = wjson.Encodestruct(v)
	return v
}
