package refund

import (
	"encoding/xml"
)

type Refundreponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg"`
	//以下字段在return_code为SUCCESS的时候有返回
	Appid                string `xml:"appid"`                          //must
	Mch_id               string `xml:"mch_id"`                         //must
	Device_info          string `xml:"device_info,omitempty"`          //maybe
	Nonce_str            string `xml:"nonce_str"`                      //must
	Sign                 string `xml:"sign"`                           //must
	Transaction_id       string `xml:"transaction_id"`                 //must
	Out_trade_no         string `xml:"out_trade_no"`                   //must
	Out_refund_no        string `xml:"out_refund_no"`                  //must
	Refund_id            string `xml:"refund_id"`                      //must
	Refund_fee           string `xml:"refund_fee"`                     //must
	Refund_fee_type      string `xml:"refund_fee_type,omitempty"`      //must
	Total_fee            string `xml:"total_fee"`                      //must
	Fee_type             string `xml:"fee_type,omitempty"`             //must
	Cash_fee             string `xml:"cash_fee"`                       //must
	Cash_fee_type        string `xml:"cash_fee_type,omitempty"`        //must
	Cash_refund_fee      string `xml:"cash_refund_fee,omitempty"`      //must
	Cash_refund_fee_type string `xml:"cash_refund_fee_type,omitempty"` //must

	// Coupon_refund_fee string `xml:"coupon_refund_fee,omitempty"` //must
	// Coupon_count string `xml:"coupon_count,omitempty"` //must
	// 	coupon_batch_id_$n  string `xml:"coupon_batch_id_$n,omitempty"`            //must
	// coupon_id_$n  string `xml:"coupon_id_$n,omitempty"`            //must
	// coupon_fee_$n  string `xml:"coupon_fee_$n,omitempty"`            //must

	Result_code  string `xml:"result_code"`
	Err_code     string `xml:"err_code,omitempty"`
	Err_code_des string `xml:"err_code_des,omitempty"`

	ReponseXML string `xml:"-"` //结果串

}
