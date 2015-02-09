package refundquery

import (
	"encoding/xml"
)

type Refundqueryresponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg"`
	//以下字段在return_code为SUCCESS的时候有返回
	Result_code  string `xml:"result_code"`
	Err_code     string `xml:"err_code,omitempty"`
	Err_code_des string `xml:"err_code_des,omitempty"`
	Appid        string `xml:"appid"`                 //must
	Mch_id       string `xml:"mch_id"`                //must
	Device_info  string `xml:"device_info,omitempty"` //maybe

	Nonce_str      string `xml:"nonce_str"` //must
	Sign           string `xml:"sign"`      //must
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Total_fee      string `xml:"total_fee"`
	Fee_type       string `xml:"fee_type,omitempty"`
	Cash_fee       string `xml:"cash_fee"`
	Cash_fee_type  string `xml:"cash_fee_type,omitempty"`

	Refund_fee string `xml:"refund_fee"`

	Coupon_refund_fee string `xml:"coupon_refund_fee,omitempty"`

	Refund_count string `xml:"refund_count"`

	ReponseXML string `xml:"-"` //结果串
}
