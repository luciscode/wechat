package imagereponse

import (
	"encoding/xml"
	"wechat/wlog"
	"wechat/wtime"
	"wechat/wxml"
)

type Imageresponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MediaId      string   `xml:"Image>MediaId"`

	ReponseXML string `xml:"-"` //结果串

}

func Getimageresponse(tousername, fromeusername, mediaid string) *Imageresponse {

	v := &Imageresponse{MsgType: "image", CreateTime: wtime.Time_stamp()}
	v.MediaId = mediaid
	v.ToUserName = tousername
	v.FromUserName = fromeusername
	v.Xml()
	return v

}
func (v *Imageresponse) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.ReponseXML = xmlresult
	wlog.PrintlnW(xmlresult, true)

	return err

}
