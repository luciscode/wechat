package kfaccount

import (
	// "auth"
	// "fmt"
	// "kfmusicmessage"

	"net/http"
	//"whttp"
)

func Sendkfmusicmessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// a := auth.Access_token(nil, nil)
	// sendkfmessageurl := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + a.Access_token
	// v := kfmusicmessage.Getkfmusicmessage("opu_njijS3dyQV6oGrGqkxyInKjM", "_Lxxli1YvskZkmqlxUw6UV-5UtlkVGLH8mszyL1X7clBsWH6ueDjDJ0t8fJoQt26")
	// fmt.Println(v.MessageJSON)
	// data := whttp.Post(sendkfmessageurl, v.MessageJSON)
	// fmt.Println(string(data))

}
