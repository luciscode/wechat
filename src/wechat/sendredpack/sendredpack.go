package sendredpack

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

func Sendredpack(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := &Sendredpackrequest{Wxappid: config.APP_ID, Mch_id: config.MCH_ID}
	request.Act_id = "RedPack-20150106"
	request.Act_name = "新年红包"

	request.Mch_billno = config.MCH_ID + "20150107" + "0000046540"
	request.Nick_name = "微信支付"
	request.Send_name = "orionhan"
	request.Re_openid = "oBmIts3-PzQq40Us1DKH74b4fjhQ"
	request.Total_amount = "1"
	request.Min_value = "1"
	request.Max_value = "1"
	request.Total_num = "1"
	request.Wishing = "恭喜发财"
	request.Client_ip = "127.0.0.1"
	request.Remark = "2015年新年红包"
	request.Logo_imgurl = "https://wx.gtimg.com/mch/img/ico-logo.png"
	request.Share_content = "分享红包内容"
	request.Share_url = "https://xx/img/wxpaylogo.png"
	request.Share_imgurl = "https://xx/img/wxpaylogo.png"
	request.Nonce_str = rand.Getnoncestr(32)

	request.Signmd5()
	request.Xml()

	fmt.Fprintf(w, request.RequestXML)

	reponse := request.DoSSLrequest()
	fmt.Println(reponse.ReponseXML)
	fmt.Println("========================================")
}
