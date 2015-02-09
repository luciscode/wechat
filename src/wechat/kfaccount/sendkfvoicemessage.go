package kfaccount

import (
	"fmt"
	"wechat/auth"
	"wechat/kfvoicemessage"

	"net/http"
	"wechat/whttp"
)

func Sendkfvoicemessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	a := auth.Access_token(nil, nil)
	sendkfmessageurl := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + a.Access_token
	v := kfvoicemessage.Getkfvoicemessage("opu_njijS3dyQV6oGrGqkxyInKjM", "RbN2ResyuLjbpLZ09CkTsIKfU8AlO_O4oRYwHd6jxMIA6dXWZBJVbDnWul-L20NV")
	fmt.Println(v.MessageJSON)
	data := whttp.Post(sendkfmessageurl, v.MessageJSON)
	fmt.Println(string(data))

}
