package closeorder

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

func Closeorder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := &Closeorderrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Nonce_str = rand.Getnoncestr(32)
	request.Out_trade_no = r.FormValue("out_trade_no")
	request.Signmd5()
	request.Xml()
	reponse := request.Dorequest()
	fmt.Println(reponse.ReponseXML)
	fmt.Println("关闭订单结果:订单" + request.Out_trade_no + "关闭结果:" + reponse.Return_msg)
	fmt.Fprintf(w, reponse.Return_msg)

	fmt.Println("========================================")

}
