package wxml

import (
	"encoding/xml"
)

//把byte数据格式化为v
func Decodebytes(b []byte, v interface{}) error {
	err := xml.Unmarshal(b, v)
	return err

}
