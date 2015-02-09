package downloadbill

import (
	"encoding/xml"
	"strings"
	"wechat/config"
	"wechat/whttp"
	"wechat/wlog"
	"wechat/wsign"
	"wechat/wxml"
)

type Downloadbillrequest struct {
	XMLName     xml.Name `xml:"xml"`
	Appid       string   `xml:"appid"`                 //公众账号ID
	Bill_date   string   `xml:"bill_date"`             //对账单日起
	Bill_type   string   `xml:"bill_type,omitempty"`   //账单类型
	Device_info string   `xml:"device_info,omitempty"` //设备号
	Mch_id      string   `xml:"mch_id"`                //商户号
	Nonce_str   string   `xml:"nonce_str"`             //随机字符串
	Sign        string   `xml:"sign"`                  //签名
	RequestXML  string   `xml:"-"`                     //最终请求串

}

func (v *Downloadbillrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult
	wlog.PrintlnW(xmlresult, false)

	return err

}

func (v *Downloadbillrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	wlog.PrintlnW(sign, false)
	return true
}
func (v Downloadbillrequest) Dorequest() Downloadbillresponse {

	data := whttp.Post(config.URL_DOWNLOADBILL, v.RequestXML)
	downloadbillresponse := Downloadbillresponse{}
	if strings.Contains(string(data), "return_code") {
		wxml.Decodebytes(data, &downloadbillresponse)
		downloadbillresponse.ReponseXML = string(data)
	} else {
		downloadbillresponse.ReponseSTR = string(data)
	}
	return downloadbillresponse

}
