package micropay

import (
	"fmt"
	"net/http"
	"wechat/config"
	"wechat/rand"
)

//1.执行刷卡支付功能入口
func Micropay(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	//获得要支付的金额
	fee := r.FormValue("fee")

	//获取用户的刷卡二维码
	authcode := r.FormValue("authcode")

	//创建Mircopayrequest请求参数
	v := &Mircopayrequest{Appid: config.APP_ID, Mch_id: config.MCH_ID}
	v.Auth_code = authcode
	if fee != "" {
		v.Total_fee = fee
	} else {
		v.Total_fee = "1"
	}
	v.Body = "微信刷卡支付"
	v.Nonce_str = rand.Getnoncestr(32)
	v.Out_trade_no = rand.Getnoncestr(32)
	v.Spbill_create_ip = "127.0.0.1"

	//对请求参数进行MD5签名，得到签名字串
	v.Signmd5()

	//把请求参数转换为xml格式的数据
	v.Xml()

	//发起支付请求，并得到返回结果
	response := v.Dorequest()

	//输出支付结果信息
	fmt.Printf("返回状态：%s \n", response.Return_code)
	fmt.Printf("返回信息: %s \n", response.Return_msg)
	fmt.Printf("业务结果: %s \n", response.Result_code)
	fmt.Printf("错误代码: %s \n", response.Err_code)
	fmt.Printf("错误代码描述: %s\n", response.Err_code_des)
	fmt.Println("========================================")

	//发送到网页
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域,解决跨域问题
	fmt.Fprintf(w, response.ReponseXML)

}
