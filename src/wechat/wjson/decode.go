package wjson
import (
	"encoding/json"
)


	
//把byte数据格式化为v
func Decodebytes(b []byte, v interface{}) error {
	err := json.Unmarshal(b, v)
	return err

}
