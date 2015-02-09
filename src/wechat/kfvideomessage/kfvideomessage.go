package kfvideomessage

import (
	"wechat/wjson"
)

func Getkfvideomessage(touser, media_id string) Kfvideomessage {
	v := Kfvideomessage{}
	v.Msgtype = "video"
	v.Touser = touser

	c := Content{}
	c.Media_id = media_id
	c.Title = "测试视频"
	v.Context = c
	v.MessageJSON, _ = wjson.Encodestruct(v)
	return v
}
