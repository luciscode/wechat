package kfaccount

import (
	"fmt"
	"wechat/auth"
	"wechat/kfimagemessage"

	"net/http"
	"wechat/whttp"
)

func Sendkfimagemessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	a := auth.Access_token(nil, nil)
	sendkfmessageurl := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + a.Access_token
	v := kfimagemessage.Getkfimagemessage("opu_njijS3dyQV6oGrGqkxyInKjM", "-hRKEbw0Pk6rm-p967HjsJrfg98V3Y-U4iNdWrDk3FjI_ZfISGJPsF5k6gsJBWx7")
	fmt.Println(v.MessageJSON)
	data := whttp.Post(sendkfmessageurl, v.MessageJSON)
	fmt.Println(string(data))

}
