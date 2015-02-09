package kfaccount

import (
	"fmt"
	"net/http"
	"wechat/auth"
	"wechat/kftextmessages"
	"wechat/whttp"
)

func Sendkftextmessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	a := auth.Access_token(nil, nil)
	sendkfmessageurl := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + a.Access_token
	v := kftextmessages.Getkftextmessage("opu_njijS3dyQV6oGrGqkxyInKjM", "testkf")
	fmt.Println(v.MessageJSON)
	data := whttp.Post(sendkfmessageurl, v.MessageJSON)
	fmt.Println(string(data))

}
