package musicresponse

import (
	"encoding/xml"
	"wechat/wlog"
	"wechat/wtime"
	"wechat/wxml"
)

type Musicresponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`

	Title        string `xml:"Music>Title,omitempty"`
	Description  string `xml:"Music>Description,omitempty"`
	MusicUrl     string `xml:"Music>MusicUrl,omitempty"`
	HQMusicUrl   string `xml:"Music>HQMusicUrl,omitempty"`
	ThumbMediaId string `xml:"Music>ThumbMediaId"`

	ReponseXML string `xml:"-"` //结果串

}

func Getmusicresponse(tousername, fromeusername, title, description, musicurl, hqmusicurl, thummediaid string) *Musicresponse {

	v := &Musicresponse{MsgType: "music", CreateTime: wtime.Time_stamp()}
	v.ToUserName = tousername
	v.FromUserName = fromeusername
	v.Title = title
	v.Description = description
	v.MusicUrl = musicurl
	v.HQMusicUrl = hqmusicurl
	v.ThumbMediaId = thummediaid
	v.Xml()
	return v

}
func (v *Musicresponse) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.ReponseXML = xmlresult
	wlog.PrintlnW(xmlresult, true)

	return err

}
