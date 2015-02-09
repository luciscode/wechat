package in

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"wechat/wxml"
)

var reponsewriter http.ResponseWriter

func Msg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	echostr := r.FormValue("echostr")
	reponsewriter = w
	if len(echostr) > 0 { //如果是绑定消息接口
		fmt.Fprintf(reponsewriter, echostr)
		// m := make(map[string]string)
		// for k, v := range r.Form {

		// 	m[k] = strings.Join(v, "")
		// 	fmt.Println("key:", k)

		// 	fmt.Println("val:", m[k])
		// }
		//checkSignature(m)
		return
	}
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(" err:", err)
	}
	//普通消息
	if strings.Contains(string(result), "<MsgId>") {
		v := Usualmessage{}
		wxml.Decodebytes(result, &v)
		err := v.insertdbusualmessage()
		if err == nil {
			fmt.Println("普通消息存储成功")
		} else {
			fmt.Println(err)

		}

		v.handleusualmessage()
	} else { //推送消息
		v := Pushmessage{}
		wxml.Decodebytes(result, &v)
		err := v.insertdbpushmessages()
		if err == nil {
			fmt.Println("推送消息存储成功")
		} else {
			fmt.Println(err)

		}
		v.handlepushmessage()

	}
	fmt.Println("==============================================")

}

func Getresponsewriter() http.ResponseWriter {
	return reponsewriter

}
func checkSignature(m map[string]string) {
	token := "wechat"

	signature := m["signature"]
	timestamp := m["timestamp"]
	nonce := m["nonce"]
	signstr := make([]string, 0)
	signstr = append(signstr, token)
	signstr = append(signstr, timestamp)
	signstr = append(signstr, nonce)
	fmt.Printf("length len=%d cap=%d %v\n", len(signstr), cap(signstr), signstr)
	a := sort.StringSlice(signstr)
	a.Sort()
	fmt.Printf("length len=%d cap=%d %v\n", len(a), cap(a), a)
	h := sha1.New()
	io.WriteString(h, a[0])
	io.WriteString(h, a[1])
	io.WriteString(h, a[2])

	fmt.Printf("% x", h.Sum(nil))
	fmt.Println(signature)

}
