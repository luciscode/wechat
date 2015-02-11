package auth

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/whttp"
	"wechat/wjson"
	// "strings"
)

type Resultauth struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Rfresh_token string `json:"rfresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func Getopenid(w http.ResponseWriter, r *http.Request) (openid string) {
	r.ParseForm()
	code := r.FormValue("code")
	requestUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + config.APP_ID + "&secret=" + config.APP_SECRET + "&code=" + code + "&grant_type=authorization_code"
	rebots := whttp.Get(requestUrl)

	var m Resultauth
	err := wjson.Decodebytes(rebots, &m)
	if err != nil {
		fmt.Println("error:", err)
	}
	openid = m.Openid

	fmt.Println("openid=" + openid)
	return openid

}
func getauthurl() string {
	str := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx2b029c08a6232582&redirect_uri=http%3a%2f%2f121.40.35.3%2fgetopenid&response_type=code&scope=snsapi_base&state=test#wechat_redirect"
	return str

}
