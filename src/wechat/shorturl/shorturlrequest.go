package shorturl

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wsign"
	"wechat/wxml"
)

type Shorturlrequest struct {
	XMLName    xml.Name `xml:"xml"`
	Appid      string   `xml:"appid"`     //must
	Long_url   string   `xml:"long_url"`  //maybe
	Mch_id     string   `xml:"mch_id"`    //maybe\
	Nonce_str  string   `xml:"nonce_str"` //must
	Sign       string   `xml:"sign"`      //must
	RequestXML string   `xml:"-"`         //最终请求串

}

func (v *Shorturlrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult

	return err

}

func (v *Shorturlrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	return true
}
func (v Shorturlrequest) Dorequest() Shorturlresponse {
	data := whttp.Post(config.URL_SHORTURL, v.RequestXML)

	shorturlresponse := Shorturlresponse{}
	wxml.Decodebytes(data, &shorturlresponse)
	shorturlresponse.ReponseXML = string(data)
	return shorturlresponse

}
