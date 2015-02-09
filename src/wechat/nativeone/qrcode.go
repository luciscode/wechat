package nativeone

import (
	"wechat/config"
	"wechat/wlog"
	"wechat/wsign"
)

//1.存储二维码信息的结构体
type Qrcode struct {
	Appid      string `xml:"appid"`      //公众帐号ID
	Mch_id     string `xml:"mch_id"`     //商户号
	Nonce_str  string `xml:"nonce_str"`  //随机字符串
	Product_id string `xml:"product_id"` //商品id，商户自定义，在用户扫码的时候会传给商户设置的回调url，商户根据该信息在自身的系统里面查找商品，然后来发起支付
	Sign       string `xml:"sign"`       //md5签名
	Time_stamp string `xml:"time_stamp"` //时间戳,毫秒级
}

//2.对Qrcode里面的有效信息进行md5签名，并存储该签名
func (v *Qrcode) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	wlog.Println(sign)
	return true
}

//3.对Qrcode里面的字段进行组装，按照微信支付要求的规则生成特定的url
func (v *Qrcode) Qrcodeurl() string {
	qrcodeurl := "weixin://wxpay/bizpayurl?sign=" + v.Sign + "&appid=" + v.Appid + "&mch_id=" + v.Mch_id + "&product_id=" + v.Product_id + "&time_stamp=" + v.Time_stamp + "&nonce_str=" + v.Nonce_str
	return qrcodeurl
}
