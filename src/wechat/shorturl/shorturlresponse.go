package shorturl

import (
	"encoding/xml"
)

type Shorturlresponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg"`
	//以下字段在return_code为SUCCESS的时候有返回

	Appid  string `xml:"appid"`  //must
	Mch_id string `xml:"mch_id"` //must

	Nonce_str   string `xml:"nonce_str"` //must
	Sign        string `xml:"sign"`      //must
	Result_code string `xml:"result_code"`
	Err_code    string `xml:"err_code,omitempty"`
	Short_url   string `xml:"short_url"`

	ReponseXML string `xml:"-"` //结果串
}
