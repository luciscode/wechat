package orderquery

import (
	"encoding/xml"
)

//存储查询订单接口的返回数据
type Orderqueryresponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg"`
	//以下字段在return_code为SUCCESS的时候有返回
	Appid        string `xml:"appid"`     //must
	Mch_id       string `xml:"mch_id"`    //must
	Nonce_str    string `xml:"nonce_str"` //must
	Sign         string `xml:"sign"`      //must
	Result_code  string `xml:"result_code"`
	Err_code     string `xml:"err_code,omitempty"`
	Err_code_des string `xml:"err_code_des,omitempty"`

	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Device_info  string `xml:"device_info,omitempty"` //maybe
	Openid       string `xml:"openid"`
	Is_subscribe string `xml:"is_subscribe"`
	Trade_type   string `xml:"trade_type"`
	Trade_state  string `xml:"trade_state"`

	Bank_type     string `xml:"bank_type"`
	Total_fee     string `xml:"total_fee"`
	Fee_type      string `xml:"fee_type,omitempty"`
	Cash_fee      string `xml:"cash_fee"`
	Cash_fee_type string `xml:"cash_fee_type,omitempty"`
	// Coupon_fee   string `xml:"coupon_fee,omitempty"`
	// Coupon_count   string `xml:"coupon_count,omitempty"`
	// Coupon_batch_id_$n   string `xml:"coupon_batch_id_$n,omitempty"`
	// Coupon_id_$n   string `xml:"coupon_id_$n,omitempty"`
	// Coupon_fee_$n   string `xml:"coupon_fee_$n,omitempty"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Attach         string `xml:"attach,omitempty"`
	Time_end       string `xml:"time_end"`

	ReponseXML string `xml:"-"` //结果串
}
