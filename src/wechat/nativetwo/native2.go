package nativetwo

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
	"wechat/unifiedorder"
)

func Native2(w http.ResponseWriter, r *http.Request) {
	v := &unifiedorder.Unifieldrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	v.Total_fee = "1"

	v.Body = "NATIVE2支付测试一分钱"
	v.Nonce_str = rand.Getnoncestr(32)
	v.Out_trade_no = rand.Getnoncestr(32)
	fmt.Println(v.Out_trade_no)
	v.Spbill_create_ip = "127.0.0.1"

	v.Notify_url = config.URL_UNIFIEDORDER_NOTIFYURL
	v.Trade_type = "NATIVE"
	v.Product_id = "native2"
	v.Signmd5() //先签名
	v.Xml()     //
	fmt.Fprintf(w, v.RequestXML)
	response := v.Dorequest()
	fmt.Fprintf(w, response.Code_url)

}
