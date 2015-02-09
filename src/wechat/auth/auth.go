package auth

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/whttp"
	"wechat/wjson"
)

var authurl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + config.APP_ID + "&secret=" + config.APP_SECRET

func Access_token(w http.ResponseWriter, r *http.Request) Accesstokenresponse {
	data := whttp.Get(authurl)
	var m Accesstokenresponse
	err := wjson.Decodebytes(data, &m)
	if err != nil {
		fmt.Println("error:", err)
	}
	return m

}
func Jsapiticket(w http.ResponseWriter, r *http.Request) Jsapiticketresponse {
	accesstoken := Access_token(nil, nil)
	jsapiticketurl := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + accesstoken.Access_token + "&type=jsapi"
	data := whttp.Get(jsapiticketurl)
	var m Jsapiticketresponse
	err := wjson.Decodebytes(data, &m)
	if err != nil {
		fmt.Println("error:", err)
	}
	return m

}
