package micropay

import (
	"encoding/xml"
	//配置文件
	"wechat/config"
	//http请求工具
	"wechat/whttp"
	//签名工具wsign
	"wechat/wsign"
	//xml格式转换工具
	"wechat/wxml"
)

//1.创建Mircopayrequest结构体，存放我们的请求字段，带有 omitempty 表示该字段选填，其他的字段属于必填
type Mircopayrequest struct {
	XMLName          xml.Name `xml:"xml"`                   //表示xml数据格式
	Appid            string   `xml:"appid"`                 //公众账号ID
	Attach           string   `xml:"attach,omitempty"`      //附加数据
	Auth_code        string   `xml:"auth_code"`             //授权码
	Body             string   `xml:"body"`                  //商品描述
	Detail           string   `xml:"detail,omitempty"`      //商品详情
	Device_info      string   `xml:"device_info,omitempty"` //设备号
	Fee_type         string   `xml:"fee_type,omitempty"`    //货币类型
	Goods_tag        string   `xml:"goods_tag,omitempty"`   //商品标记
	Mch_id           string   `xml:"mch_id"`                //商户号
	Nonce_str        string   `xml:"nonce_str"`             //随机字符串
	Out_trade_no     string   `xml:"out_trade_no"`          //商户订单号
	Sign             string   `xml:"sign"`                  //签名
	Spbill_create_ip string   `xml:"spbill_create_ip"`      //终端IP
	Time_expite      string   `xml:"time_expite,omitempty"` //交易失效时间
	Time_start       string   `xml:"time_start,omitempty"`  //交易起始时间
	Total_fee        string   `xml:"total_fee"`             //总金额

	RequestXML string `xml:"-"` //存放最终生成的xml请求数据,不参与签名以及xml的组成
}

//2.对Mircopayrequest里面的字段进行md5签名，存储到Mircopayrequest里面的Sign变量
func (v *Mircopayrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	return true
}

//3.生成Mircopayrequest对应的xml数据，存储到Mircopayrequest里面的RequestXML变量
func (v *Mircopayrequest) Xml() error {
	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult
	return err
}

//4.向刷卡接口发送post的请求，请求的数据为我们的生成存储在RequestXML变量的xml数据
//  并且将得到的数据解析到Mircopayresponse结构体
func (v Mircopayrequest) Dorequest() Mircopayresponse {
	//发送post请求,f返回数据为data
	data := whttp.Post(config.URL_MICROPAY, v.RequestXML)

	//解析data到Mircopayresponse结构体
	mircopayresponse := Mircopayresponse{}
	wxml.Decodebytes(data, &mircopayresponse)
	mircopayresponse.ReponseXML = string(data)
	//返回结果
	return mircopayresponse

}
