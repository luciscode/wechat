package downloadbill

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

func Downloadbill(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := &Downloadbillrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Nonce_str = rand.Getnoncestr(32)
	request.Bill_date = "20141206"
	request.Bill_type = "ALL"
	request.Signmd5()
	request.Xml()
	reponse := request.Dorequest()
	if reponse.Return_code == "" {
		fmt.Println("对账单查询:开始日期" + request.Bill_date + "查询结果:" + reponse.ReponseSTR)

	} else {

		fmt.Println("对账单查询异常:" + reponse.Return_msg)
	}
	fmt.Println("========================================")

}
