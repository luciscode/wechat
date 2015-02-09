package wsign

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

//1.得到v里面的字段的md5签名，v在创建的时候就要保证里面的字段是按照ascii从小到大的顺序填写
func SignMD5(v interface{}, apikey string) (string, bool) {
	var signstr bytes.Buffer

	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		name := field.Name

		keytemp := field.Tag.Get("xml")
		keymap := strings.Split(keytemp, ",")
		key := keymap[0]

		value := (reflect.Indirect(vv).FieldByName(name)).String()
		if value != "" && key != "xml" {
			signstr.WriteString(key + "=" + value + "&")

		}
	}
	signstr.WriteString("key=" + apikey)
	hasher := md5.New()
	hasher.Write([]byte(signstr.String()))
	sign := hex.EncodeToString(hasher.Sum(nil))
	sign = strings.ToUpper(sign)
	return sign, true

}
func SignSHA1(v interface{}) (string, bool) {
	var signstr bytes.Buffer

	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		name := field.Name

		keytemp := field.Tag.Get("json")
		keymap := strings.Split(keytemp, ",")
		key := keymap[0]

		value := (reflect.Indirect(vv).FieldByName(name)).String()
		if value != "" {
			signstr.WriteString(key + "=" + value)

		}

		if i < (vt.NumField() - 1) {
			signstr.WriteString("&")
		}
	}

	fmt.Println(signstr.String())
	hasher := sha1.New()
	hasher.Write([]byte(signstr.String()))
	sign := hex.EncodeToString(hasher.Sum(nil))
	return sign, true

}
