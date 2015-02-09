package jsapipay

import (
	"encoding/xml"
	"wechat/config"
	"wechat/wsign"
)

//1.定义Jsapirequest结构体，存储getBrandWCPayRequest接口要使用的参数
type Jsapirequest struct {
	XMLName   xml.Name `xml:"xml"`
	AppId     string   `xml:"appId"`     //公众账号ID
	NonceStr  string   `xml:"nonceStr"`  //随机字符串
	Package   string   `xml:"package"`   //账单类型
	SignType  string   `xml:"signType"`  //商户号
	TimeStamp string   `xml:"timeStamp"` //对账单日起

	PaySign string `xml:"-"` //最终请求串

}

//2.对Jsapirequest的非空字段进行md5签名，并存储签名结构
func (v *Jsapirequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.PaySign = sign
	return true
}
