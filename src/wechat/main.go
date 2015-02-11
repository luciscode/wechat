package main

import (
	"fmt"
	"log"
	"net/http"
	"wechat/auth"
	"wechat/card"
	"wechat/closeorder"
	"wechat/downloadbill"
	"wechat/fileupload"
	"wechat/in"
	"wechat/jsapipay"
	"wechat/jswechat"
	"wechat/kfaccount"
	"wechat/mgo"
	"wechat/micropay"
	"wechat/nativeone"
	"wechat/nativetwo"
	"wechat/notify"
	"wechat/orderquery"
	"wechat/refund"
	"wechat/refundquery"
	"wechat/sendredpack"
	"wechat/shorturl"
	"wechat/test"
)

func rout(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	path := r.URL.Path

	if path == "/message" {
		fmt.Println("有微信消息到来")
		in.Msg(w, r)
	} else if path == "/micropay" {
		fmt.Println("被扫支付")
		micropay.Micropay(w, r)

	} else if path == "/index" {
		fmt.Println("欢迎访问")
		fmt.Fprintf(w, "欢迎访问")

	} else if path == "/notify" {
		fmt.Println("支付结果异步通知")
		notify.Notify(w, r)
	} else if path == "/jsapi" {
		fmt.Println("JSAPI支付")
		jsapipay.Jsapi(w, r)
	} else if path == "/getopenid" {
		fmt.Println("网页获取openid")
		jsapipay.Getopenid(w, r)
	} else if path == "/native1" {
		fmt.Println("原生支付URL获取")
		nativeone.Native1(w, r)
	} else if path == "/nativeonecallback" {
		fmt.Println("原生支付模式回调地址")
		nativeone.Nativeonecallback(w, r)
	} else if path == "/native2" {
		fmt.Println("原生支付模式二获取支付二维码地址")
		nativetwo.Native2(w, r)
	} else if path == "/orderquery" {
		fmt.Println("查询订单")
		orderquery.Orderquery(w, r)
	} else if path == "/closeorder" {
		fmt.Println("关闭订单")
		closeorder.Closeorder(w, r)
	} else if path == "/refund" {
		fmt.Println("退款操作")
		refund.Refund(w, r)
	} else if path == "/refundquery" {
		fmt.Println("退款查询操作")
		refundquery.Refundquery(w, r)
	} else if path == "/downloadbill" {
		fmt.Println("对账单下载操作")
		downloadbill.Downloadbill(w, r)
	} else if path == "/shorturl" {
		fmt.Println("对账单下载操作")
		shorturl.Shorturl(w, r)
	} else if path == "/sendredpack" {
		fmt.Println("发送红包")
		sendredpack.Sendredpack(w, r)
	} else if path == "/accesstoken" {
		fmt.Println("获取accesstoken")
		auth.Access_token(w, r)
	} else if path == "/kfaccount/addkfaccount" {
		fmt.Println("增加客服 ")
		kfaccount.Addkfaccount(w, r)
	} else if path == "/kfaccount/sendkftextmessage" {
		fmt.Println("发送text客服消息")
		kfaccount.Sendkftextmessage(w, r)
	} else if path == "/kfaccount/sendkfimagemessage" {
		fmt.Println("发送image客服消息")
		kfaccount.Sendkfimagemessage(w, r)
	} else if path == "/kfaccount/sendkfvoicemessage" {
		fmt.Println("发送voice客服消息")
		kfaccount.Sendkfvoicemessage(w, r)
	} else if path == "/kfaccount/sendkfvideomessage" {
		fmt.Println("发送video客服消息")
		kfaccount.Sendkfvideomessage(w, r)
	} else if path == "/kfaccount/sendkfmusicmessage" {
		fmt.Println("发送music客服消息")
		kfaccount.Sendkfmusicmessage(w, r)
	} else if path == "/share" {
		fmt.Println("分享页面")
		jswechat.Share(w, r)
	} else if path == "/closeordertest" {
		fmt.Println("测试关闭订单")
		test.Closeordertest(w, r)
	} else if path == "/mongoinsertkeyword" {
		fmt.Println("插入关键字")
		mgo.MongoinsertKeyword(w, r)
	} else if path == "/mongofindkeyword" {
		fmt.Println("查找关键字")
		mgo.MongoFindKeyword(w, r)
	} else if path == "/fileupload" {
		fmt.Println("上传文件")
		fileupload.Fileupload(w, r)
	} else if path == "/createcard" {
		fmt.Println("创建卡券")
		card.Createcard(w, r)
	} else if path == "/createshakecard" {
		fmt.Println("创建摇一摇卡券")
		card.Createshakecard(w, r)
	} else if path == "/updateshakecard" {
		fmt.Println("更新摇一摇")
		card.Updateshakecard(w, r)
	} else if path == "/testpost" {
		fmt.Println("post和get测试")
		test.Testpost(w, r)
	}

}

func main() {

	fmt.Println("启动服务，创建数据库")

	http.HandleFunc("/", rout) //设置访问的路由

	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
