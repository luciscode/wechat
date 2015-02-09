package voiceresponse

import (
	"encoding/xml"
	"wechat/wlog"
	"wechat/wtime"
	"wechat/wxml"
)

type Voiceresponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MediaId      string   `xml:"Voice>MediaId"`

	ReponseXML string `xml:"-"` //结果串

}

func Getvoiceresponse(tousername, fromeusername, mediaid string) *Voiceresponse {

	v := &Voiceresponse{MsgType: "voice", CreateTime: wtime.Time_stamp()}
	v.MediaId = mediaid
	v.ToUserName = tousername
	v.FromUserName = fromeusername
	v.Xml()
	return v

}
func (v *Voiceresponse) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.ReponseXML = xmlresult
	wlog.PrintlnW(xmlresult, true)

	return err

}
