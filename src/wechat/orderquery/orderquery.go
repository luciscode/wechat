package orderquery

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

func Orderquery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := &Orderqueryrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Nonce_str = rand.Getnoncestr(32)
	request.Out_trade_no = r.FormValue("out_trade_no")
	request.Signmd5()
	request.Xml()
	reponse := request.Dorequest()
	fmt.Println("查询结果:订单" + request.Out_trade_no + "交易结果:" + reponse.Trade_state)
	fmt.Println("========================================")

}
