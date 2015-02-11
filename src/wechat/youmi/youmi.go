package youmi

import (
	"fmt"
	"net/http"
)

func Youmipay(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("code") != "" {
		openid := auth.Getopenid(w, r)
	} else {
		fmt.Println("error no openid")
	}
	fmt.Println("========================================")

}
func getauthurl() string {
	str := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx2b029c08a6232582&redirect_uri=http%3a%2f%2f121.40.35.3%2fgetopenid&response_type=code&scope=snsapi_base&state=test#wechat_redirect"
	return str

}
