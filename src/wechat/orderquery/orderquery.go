package orderquery

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

//查询订单入口
func Orderquery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//查询订单请求参数创建
	request := &Orderqueryrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Nonce_str = rand.Getnoncestr(32)
	request.Out_trade_no = r.FormValue("out_trade_no")
	request.Signmd5()
	request.Xml()
	//进行订单查询，并存储返回结果
	reponse := request.Dorequest()
	fmt.Println("查询结果:订单" + request.Out_trade_no + "交易结果:" + reponse.Trade_state)
	fmt.Println("========================================")

}
