package wjson

import (
	"encoding/json"
)

//把struct数据解压出来转化为json
func Encodestruct(v interface{}) (jsonresult string, err error) {
	output, err := json.Marshal(v)

	jsonresult = string(output)

	return jsonresult, err
}
