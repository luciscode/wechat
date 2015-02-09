package refundquery

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wsign"
	"wechat/wxml"
)

type Refundqueryrequest struct {
	XMLName        xml.Name `xml:"xml"`
	Appid          string   `xml:"appid"`                    //must
	Device_info    string   `xml:"device_info,omitempty"`    //maybe
	Mch_id         string   `xml:"mch_id"`                   //maybe\
	Nonce_str      string   `xml:"nonce_str"`                //must
	Out_refund_no  string   `xml:"out_refund_no,omitempty"`  //maybe
	Out_trade_no   string   `xml:"out_trade_no"`             //maybe
	Refund_id      string   `xml:"refund_id,omitempty"`      //maybe
	Sign           string   `xml:"sign"`                     //must
	Transaction_id string   `xml:"transaction_id,omitempty"` //maybe

	RequestXML string `xml:"-"` //最终请求串

}

func (v *Refundqueryrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult

	return err

}

func (v *Refundqueryrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	return true
}
func (v Refundqueryrequest) Dorequest() Refundqueryresponse {
	data := whttp.Post(config.URL_REFUNDQUERY, v.RequestXML)

	refundqueryresponse := Refundqueryresponse{}
	wxml.Decodebytes(data, &refundqueryresponse)
	refundqueryresponse.ReponseXML = string(data)
	return refundqueryresponse

}
