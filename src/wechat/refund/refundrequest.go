package refund

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wlog"
	"wechat/wsign"
	"wechat/wxml"
)

type Refundrequest struct {
	XMLName         xml.Name `xml:"xml"`
	Appid           string   `xml:"appid"`                     //must
	Device_info     string   `xml:"device_info,omitempty"`     //maybe
	Mch_id          string   `xml:"mch_id"`                    //must
	Nonce_str       string   `xml:"nonce_str"`                 //must
	Op_user_id      string   `xml:"op_user_id"`                //maybe
	Out_refund_no   string   `xml:"out_refund_no"`             //must
	Out_trade_no    string   `xml:"out_trade_no"`              //must
	Refund_fee      string   `xml:"refund_fee"`                //must
	Refund_fee_type string   `xml:"refund_fee_type,omitempty"` //maybe
	Sign            string   `xml:"sign"`                      //must
	Total_fee       string   `xml:"total_fee"`                 //must
	Transaction_id  string   `xml:"transaction_id"`            //maybe
	RequestXML      string   `xml:"-"`                         //最终请求串

}

func (v *Refundrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult
	wlog.PrintlnW(xmlresult, false)

	return err

}

func (v *Refundrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	wlog.PrintlnW(sign, false)
	return true
}
func (v Refundrequest) DoSSLrequest() Refundreponse {

	data := whttp.SSLPost(config.URL_REFUND, v.RequestXML)
	wlog.PrintlnW(string(data), true)
	refundreponse := Refundreponse{}
	wxml.Decodebytes(data, &refundreponse)
	refundreponse.ReponseXML = string(data)

	return refundreponse

}
