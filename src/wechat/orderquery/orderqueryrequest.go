package orderquery

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wsign"
	"wechat/wxml"
)

//1.定义结构体，存储查询订单号使用的参数,注意：out_trade_no和transaction_id必须填一个，2个都存在的情况下，优先以transaction_id为准
type Orderqueryrequest struct {
	XMLName        xml.Name `xml:"xml"`
	Appid          string   `xml:"appid"`                    //公众账号ID
	Mch_id         string   `xml:"mch_id"`                   //mch_id
	Nonce_str      string   `xml:"nonce_str"`                //随机字符串
	Out_trade_no   string   `xml:"out_trade_no,omitempty"`   //商户订单号
	Sign           string   `xml:"sign"`                     //签名
	Transaction_id string   `xml:"transaction_id,omitempty"` //微信订单号

	RequestXML string `xml:"-"` //最终请求串

}

//2.对请求的字段进行md5签名，并存储
func (v *Orderqueryrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	return true
}

//3.把Orderqueryrequest里面的非空字段生成xml字符串，并存储
func (v *Orderqueryrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult

	return err

}

//4.将xml字符串发送给查询订单接口，并把结构解析到Orderqueryresponse结构体
func (v Orderqueryrequest) Dorequest() Orderqueryresponse {
	data := whttp.Post(config.URL_ORDERQUERY, v.RequestXML)

	orderqueryresponse := Orderqueryresponse{}
	wxml.Decodebytes(data, &orderqueryresponse)
	orderqueryresponse.ReponseXML = string(data)

	return orderqueryresponse

}
