package unifiedorder

import (
	"encoding/xml"
)

//1.存储统一下单的返回的数据
type Unifiedorderreponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"` //返回状态码
	Return_msg  string   `xml:"return_msg"`  //返回信息
	//以下字段在return_code为SUCCESS的时候有返回
	Appid        string `xml:"appid"`                  //公众账号ID
	Mch_id       string `xml:"mch_id"`                 //商户号
	Device_info  string `xml:"device_info,omitempty"`  //设备号
	Nonce_str    string `xml:"nonce_str"`              //随机字符串
	Sign         string `xml:"sign"`                   //签名
	Result_code  string `xml:"result_code"`            //业务结果
	Err_code     string `xml:"err_code,omitempty"`     //错误代码
	Err_code_des string `xml:"err_code_des,omitempty"` //错误代码描述
	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Trade_type string `xml:"trade_type,omitempty"` //交易类型
	Prepay_id  string `xml:"prepay_id,omitempty"`  //预支付交易会话标识
	Code_url   string `xml:"code_url,omitempty"`   //二维码链接

	ReponseXML string `xml:"-"` //存储返回的xml串

}
