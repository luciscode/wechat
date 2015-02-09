package unifiedorder

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wlog"
	"wechat/wsign"
	"wechat/wxml"
)

//1.创建Unifieldrequest结构体，存放要使用的数据，omitempty 表示该字段选填，否则为必填
type Unifieldrequest struct {
	XMLName          xml.Name `xml:"xml"`
	Appid            string   `xml:"appid"`                 //公众账号ID
	Attach           string   `xml:"attach,omitempty"`      //附加数据
	Body             string   `xml:"body"`                  //商品描述
	Detail           string   `xml:"detail,omitempty"`      //商品详情
	Device_info      string   `xml:"device_info,omitempty"` //设备号
	Fee_type         string   `xml:"fee_type,omitempty"`    //货币类型
	Notify_url       string   `xml:"notify_url"`            //通知地址
	Goods_tag        string   `xml:"goods_tag,omitempty"`   //商品标记
	Mch_id           string   `xml:"mch_id"`                //商户号
	Nonce_str        string   `xml:"nonce_str"`             //随机字符串
	Openid           string   `xml:"openid,omitempty"`      //用户标识
	Out_trade_no     string   `xml:"out_trade_no"`          //商户订单号
	Product_id       string   `xml:"product_id,omitempty"`  //商品ID
	Sign             string   `xml:"sign"`                  //签名
	Spbill_create_ip string   `xml:"spbill_create_ip"`      //终端IP
	Time_expire      string   `xml:"time_expire,omitempty"` //交易结束时间
	Time_start       string   `xml:"time_start,omitempty"`  //交易起始时间
	Total_fee        string   `xml:"total_fee"`             //总金额
	Trade_type       string   `xml:"trade_type"`            //交易类型

	RequestXML string `xml:"-"` //存放最终请求xml串

}

//2.对Unifieldrequest里面的非空字段进行MD5签名，得到签名结构，并保存该结果
func (v *Unifieldrequest) Signmd5() bool {
	signmd5, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = signmd5
	wlog.PrintlnW(signmd5, false)
	return true
}

//3.将Unifieldrequest里面的非空字段组织为xml格式的数据
func (v *Unifieldrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult
	wlog.PrintlnW(xmlresult, false)
	return err

}

//4.将最终生成的xml数据发送到统一支付接口，并把返回的结果解析到Unifiedorderreponse
func (v Unifieldrequest) Dorequest() Unifiedorderreponse {

	data := whttp.Post(config.URL_UNIFIEDORDER, v.RequestXML)
	unifiedorderreponse := Unifiedorderreponse{}
	wxml.Decodebytes(data, &unifiedorderreponse)
	unifiedorderreponse.ReponseXML = string(data)
	return unifiedorderreponse

}
