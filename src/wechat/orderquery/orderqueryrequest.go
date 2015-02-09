package orderquery

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wsign"
	"wechat/wxml"
)

type Orderqueryrequest struct {
	XMLName        xml.Name `xml:"xml"`
	Appid          string   `xml:"appid"`                    //must
	Mch_id         string   `xml:"mch_id"`                   //maybe\
	Nonce_str      string   `xml:"nonce_str"`                //must
	Out_trade_no   string   `xml:"out_trade_no,omitempty"`   //maybe
	Sign           string   `xml:"sign"`                     //must
	Transaction_id string   `xml:"transaction_id,omitempty"` //maybe
	RequestXML     string   `xml:"-"`                        //最终请求串

}

func (v *Orderqueryrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult

	return err

}

func (v *Orderqueryrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	return true
}
func (v Orderqueryrequest) Dorequest() Orderqueryresponse {
	data := whttp.Post(config.URL_ORDERQUERY, v.RequestXML)

	orderqueryresponse := Orderqueryresponse{}
	wxml.Decodebytes(data, &orderqueryresponse)
	orderqueryresponse.ReponseXML = string(data)

	return orderqueryresponse

}
