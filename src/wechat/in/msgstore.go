package in

import (
	// "bytes"
	// "fmt"
	"reflect"
	"strings"
	"wechat/mgo"
)

func (v Usualmessage) insertdbusualmessage() error {
	usual := map[string]string{}

	// err := mgo.InsertUsualmessage(v)

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
			usual[key] = value
		}
	}
	err := mgo.MongoInsert(usual, "usual")

	return err

}
func (v Pushmessage) insertdbpushmessages() error {
	push := map[string]string{}

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
			push[key] = value
		}
	}
	err := mgo.MongoInsert(push, "push")
	return err

}
