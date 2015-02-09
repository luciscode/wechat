package jswechat

import (
	"fmt"
	"html/template"
	"net/http"
	"wechat/auth"
	"wechat/config"
	"wechat/rand"
	"wechat/wsign"
	"wechat/wtime"
)

func Share(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	t, err := template.ParseFiles("template/share.html")
	fmt.Println(err)

	da := Config{}

	da.AppId = config.APP_ID
	da.Debug = true
	da.Timestamp = wtime.Time_stamp()
	da.NonceStr = rand.Getnoncestr(10)
	da.JsApiList = "onMenuShareTimeline"

	si := Sha1message{}
	si.NonceStr = da.NonceStr
	si.Timestamp = da.Timestamp
	jsticket := auth.Jsapiticket(nil, nil)
	si.Jsapi_ticket = jsticket.Ticket
	si.Url = "http://121.40.35.3/jswechat/share"
	da.Signature, _ = wsign.SignSHA1(si)
	fmt.Println(da.Signature)
	t.Execute(w, da)
}
