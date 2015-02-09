package micropay

import (
	"encoding/xml"
)

//1.创建Mircopayresponse结构体，存放我们的请求字段，带有 omitempty 表示该字段选填，其他的字段属于必填

type Mircopayresponse struct {
	XMLName        xml.Name `xml:"xml"`
	Return_code    string   `xml:"return_code"`           //返回状态码
	Return_msg     string   `xml:"return_msg"`            //返回信息
	Appid          string   `xml:"appid"`                 //公众账号ID
	Mch_id         string   `xml:"mch_id"`                //商户号
	Device_info    string   `xml:"device_info,omitempty"` //设备号
	Nonce_str      string   `xml:"nonce_str"`             //随机字符串
	Sign           string   `xml:"sign"`                  //签名
	Result_code    string   `xml:"result_code"`           //业务结果
	Err_code       string   `xml:"err_code"`              //错误代码
	Err_code_des   string   `xml:"err_code_des"`          //错误代码描述
	Is_subscribe   string   `xml:"is_subscribe"`          //是否关注公众账号
	Trade_type     string   `xml:"trade_type"`            //交易类型
	Bank_type      string   `xml:"bank_type"`             //付款银行
	Fee_type       string   `xml:"fee_type"`              //货币类型
	Total_fee      string   `xml:"total_fee"`             //总金额
	Cash_fee_type  string   `xml:"cash_fee_type"`         //现金支付货币类型
	Cash_fee       string   `xml:"cash_fee"`              //现金支付金额
	Coupon_fee     string   `xml:"coupon_fee"`            //代金券或立减优惠金额
	Transaction_id string   `xml:"transaction_id"`        //微信支付订单号
	Out_trade_no   string   `xml:"out_trade_no"`          //商户订单号
	Attach         string   `xml:"attach"`                //商家数据包
	Time_end       string   `xml:"time_end"`              //支付完成时间

	ReponseXML string `xml:"-"` //结果xml串

}
