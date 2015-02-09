package kfmusicmessage

import (
	"wechat/wjson"
)

func Getkfmusicmessage(touser, thumb_media_id string) Kfmucismessage {
	v := Kfmucismessage{}
	v.Msgtype = "music"
	v.Touser = touser

	c := Content{}
	c.Thumb_media_id = thumb_media_id
	c.Title = "测试视频"
	v.Context = c
	v.MessageJSON, _ = wjson.Encodestruct(v)
	return v
}
