package notify

import (
	"fmt"
	"net/http"
	"wechat/whttp"
	"wechat/wxml"
)

func Notify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := whttp.Parserequest(r)
	notifyresult := &Notifyresult{}
	wxml.Decodebytes(data, notifyresult)
	notifyresult.ReponseXML = string(data)
	fmt.Println(r.Method)
	fmt.Println("完成支付，微信订单号为:" + notifyresult.Transaction_id)

	pb := Notifyrequest{Return_code: notifyresult.Return_code, Return_msg: notifyresult.Return_msg}
	pb.Xml()
	fmt.Fprintf(w, pb.RequestXML)

}
