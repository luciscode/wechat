package shorturl

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

func Shorturl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := &Shorturlrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Nonce_str = rand.Getnoncestr(32)
	request.Long_url = "weixin://wxpay/bizpayurl?sign=0BACF38077A3E26C0202E1C5599E1DCA&appid=wx2b029c08a6232582&mch_id=10021642&product_id=dsad&time_stamp=1420562359&nonce_str=jOHbaK5JHPRZGimsaeWSeFLb3Or8Q9Xu"
	request.Signmd5()
	request.Xml()
	reponse := request.Dorequest()
	fmt.Println("转换后的短链接为:" + reponse.Short_url)
	fmt.Fprintf(w, "转换后的短链接为:"+reponse.Short_url)

	fmt.Println("========================================")

}
