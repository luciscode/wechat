package textreponse

import (
	"encoding/xml"
	"wechat/wlog"
	"wechat/wtime"
	"wechat/wxml"
)

type Textresponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`

	ReponseXML string `xml:"-"` //结果串

}

func Gettextresponse(tousername, fromeusername, content string) *Textresponse {

	v := &Textresponse{MsgType: "text", CreateTime: wtime.Time_stamp()}
	v.Content = content
	v.ToUserName = tousername
	v.FromUserName = fromeusername
	v.Xml()
	return v

}
func (v *Textresponse) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.ReponseXML = xmlresult
	wlog.PrintlnW(xmlresult, true)

	return err

}
