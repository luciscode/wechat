package refund

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

func Refund(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := &Refundrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Nonce_str = rand.Getnoncestr(32)

	request.Transaction_id = r.FormValue("transaction_id")
	request.Out_trade_no = r.FormValue("out_trade_no")
	request.Out_refund_no = rand.Getnoncestr(32)
	request.Total_fee = "1"
	request.Refund_fee = "1"
	request.Op_user_id = config.MCH_ID

	request.Signmd5()
	request.Xml()
	v := request.DoSSLrequest()
	fmt.Println(v.Err_code_des)

	fmt.Println("========================================")

}
