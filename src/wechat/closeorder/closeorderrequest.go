package closeorder

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wsign"
	"wechat/wxml"
)

type Closeorderrequest struct {
	XMLName      xml.Name `xml:"xml"`
	Appid        string   `xml:"appid"`                  //must
	Mch_id       string   `xml:"mch_id"`                 //maybe\
	Nonce_str    string   `xml:"nonce_str"`              //must
	Out_trade_no string   `xml:"out_trade_no,omitempty"` //maybe
	Sign         string   `xml:"sign"`                   //must
	RequestXML   string   `xml:"-"`                      //最终请求串
}

func (v *Closeorderrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult

	return err

}

func (v *Closeorderrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	return true
}
func (v Closeorderrequest) Dorequest() Closeorderresponse {
	response := whttp.Post(config.URL_CLOSEORDER, v.RequestXML)

	closeorderresponse := Closeorderresponse{}
	wxml.Decodebytes(response, &closeorderresponse)
	closeorderresponse.ReponseXML = string(response)
	return closeorderresponse

}
