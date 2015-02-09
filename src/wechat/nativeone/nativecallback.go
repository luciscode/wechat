package nativeone

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"wechat/config"
	"wechat/rand"
	"wechat/unifiedorder"
	"wechat/wxml"
)

//1.存储用户扫码二维码时候，微信支付发送给回调url的数据
type Callback struct {
	XMLName      xml.Name `xml:"xml"`
	Appid        string   `xml:"appid"`        //公众号ID
	Is_subscribe string   `xml:"is_subscribe"` //标识用户是否关注该公众号
	Mch_id       string   `xml:"mch_id"`       //商户号
	Nonce_str    string   `xml:"nonce_str"`    //随机字符串
	Openid       string   `xml:"openid"`       //用户openid
	Product_id   string   `xml:"product_id"`   //商品id
	Sign         string   `xml:"sign"`         //签名
}

//1.回调url，实现统一下单操作，并把下单结果发送给微信支付后台
func Nativeonecallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//将微信支付后台发送的二维码相关信息解析存储到Callback结构体
	result, _ := ioutil.ReadAll(r.Body)
	vw := &Callback{}
	wxml.Decodebytes(result, vw)

	//进行统一下单
	v := &unifiedorder.Unifieldrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	v.Total_fee = "1"
	v.Body = "NATIVE1支付测试一分钱"
	v.Nonce_str = rand.Getnoncestr(32)
	v.Out_trade_no = rand.Getnoncestr(32)
	v.Spbill_create_ip = "127.0.0.1"
	v.Notify_url = config.URL_UNIFIEDORDER_NOTIFYURL
	v.Trade_type = "NATIVE"
	v.Product_id = vw.Product_id
	v.Signmd5()
	v.Xml()
	response := v.Dorequest()

	//将统一下单的结果发送给微信支付后台
	xmlresult, _ := wxml.Endoestruct(response)
	fmt.Fprintf(w, xmlresult)

	fmt.Println("========================================")

}
