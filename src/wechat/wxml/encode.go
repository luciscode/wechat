package wxml

import (
	"encoding/xml"
)

//把结构体v 转化为xml字符串，并返回
func Endoestruct(v interface{}) (xmlresult string, err error) {
	output, err := xml.MarshalIndent(v, "  ", "    ")

	xmlresult = string(output)

	return xmlresult, err
}
