package downloadbill

import (
	"encoding/xml"
)

type Downloadbillresponse struct {
	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"` //返回状态码
	Return_msg  string   `xml:"return_msg"`  //返回信息

	ReponseXML string `xml:"-"` //下载失败时候的结果xml串
	ReponseSTR string `xml:"-"` //下载成功时候的账单字符串
}
