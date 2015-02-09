package notify

import (
	"encoding/xml"
	"wechat/wxml"
)

type Notifyrequest struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg"`
	RequestXML  string   `xml:"-"` //最终请求串

}

func (v *Notifyrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult

	return err

}
