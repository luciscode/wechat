package wmap

import (
	"reflect"
	"strings"
)

func Struct2map(v interface{}) map[string]string {
	datamap := map[string]string{}

	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		name := field.Name

		keytemp := field.Tag.Get("xml")
		keymap := strings.Split(keytemp, ",")
		key := keymap[0]

		value := (reflect.Indirect(vv).FieldByName(name)).String()
		if value != "" {
			datamap[key] = value
		}
	}
	return datamap
}
