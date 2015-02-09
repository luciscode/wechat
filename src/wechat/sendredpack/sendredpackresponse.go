package sendredpack

import (
	"encoding/xml"
)

type Sendredpackreponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg,,omitempty"`
	//以下字段在return_code为SUCCESS的时候有返回
	Sign         string `xml:"sign"` //must
	Result_code  string `xml:"result_code"`
	Err_code     string `xml:"err_code,omitempty"`
	Err_code_des string `xml:"err_code_des,omitempty"`
	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	mch_billno string `xml:"mch_billno"` //must
	Mch_id     string `xml:"mch_id"`     //must
	wxappid    string `xml:"wxappid"`    //maybe
	re_openid  string `xml:"re_openid"`  //must

	total_amount string `xml:"total_amount"` //must

	ReponseXML string `xml:"-"` //结果串

}
