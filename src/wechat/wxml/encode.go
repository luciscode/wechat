package wxml

import (
	"encoding/xml"
)

//把struct数据解压出来转化为xml
func Endoestruct(v interface{}) (xmlresult string, err error) {
	output, err := xml.MarshalIndent(v, "  ", "    ")

	xmlresult = string(output)

	return xmlresult, err
}
