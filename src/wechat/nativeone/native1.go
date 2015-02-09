package nativeone

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"wechat/config"
	"wechat/rand"
)

//获取模式1二维码，打印到网页
func Native1(w http.ResponseWriter, r *http.Request) {

	v := &Qrcode{Appid: config.APP_ID, Mch_id: config.MCH_ID, Time_stamp: strconv.FormatInt(time.Now().Unix(), 10), Nonce_str: rand.Getnoncestr(32), Product_id: "dsad"}
	v.Signmd5()

	fmt.Fprintf(w, v.Qrcodeurl())

}
