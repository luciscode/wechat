package jsapipay

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"wechat/config"
	"wechat/rand"
	"wechat/unifiedorder"
	"wechat/wlog"
	"wechat/wtime"
)

func Jsapi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// code := r.FormValue("code")
	// if code == "" {
	// 	http.Redirect(w, r, getjsapiurl(), http.StatusForbidden)
	// }

	//1.统一下单
	v := &unifiedorder.Unifieldrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}

	//支付金额，单位为分
	v.Total_fee = "1"
	//商品说明
	v.Body = "JSAPI支付测试一分钱"
	//32位随机字符串
	v.Nonce_str = rand.Getnoncestr(32)
	//商户订单号，此处也随机生成
	v.Out_trade_no = rand.Getnoncestr(32)
	//发起支付的机器IP
	v.Spbill_create_ip = "127.0.0.1"
	//设置openid,可以写死也可以根据code获取
	if r.FormValue("code") != "" {
		v.Openid = Getopenid(w, r)
	} else {
		v.Openid = "omL67jm0A1sKwystTq7WsU28MF_c"
	}
	//最终支付成功的通知地址
	v.Notify_url = config.URL_UNIFIEDORDER_NOTIFYURL
	//支付方式为JSAPI
	v.Trade_type = "JSAPI"
	//对上面设置的字段进行签名
	v.Signmd5()
	//把所有的有效字段组织为xml格式
	v.Xml()

	//向统一下单接口发起请求，并把返回请求结果
	res := v.Dorequest()
	fmt.Println("prepayid=", res.Prepay_id)
	fmt.Println("========================================")

	//打开支付网页,并传递参数,包括传递我们上面统一下单获取到的prepay_id
	if r.Method == "GET" {
		path, _ := os.Getwd()
		t, err := template.ParseFiles(path + "/jsapi.html")
		fmt.Println(err)
		da := Jsapirequest{AppId: config.APP_ID}
		da.Package = "prepay_id=" + res.Prepay_id
		da.TimeStamp = wtime.Time_stamp()
		da.NonceStr = rand.Getnoncestr(32)
		da.SignType = "MD5"
		da.Signmd5()

		t.Execute(w, da)
	}
}

func Callback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(" err:", err)
	} else {
		wlog.Println(string(result))
	}

}
func getjsapiurl() string {
	str := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx2b029c08a6232582&redirect_uri=http%3a%2f%2f121.40.35.3%2fjsapi&response_type=code&scope=snsapi_base&state=test#wechat_redirect"
	return str

}
