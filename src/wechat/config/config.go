package config

//微信支付信息

//debug
// var APP_ID = "****************"
// var APP_SECRET = "*********************"
var APP_ID = "wx2b029c08a6232582"
var APP_SECRET = "fc062a51903db580d8a5dc87d3fad0ef"
var MCH_ID = "10021642"
var SUB_MCH_ID = ""
var API_KEY = "134acc6ada2e13ab39038d257f47c117"

var LOG_ON = false

var URL_UNIFIEDORDER = "https://api.mch.weixin.qq.com/pay/unifiedorder"
var URL_UNIFIEDORDER_NOTIFYURL = "http://121.40.35.3/notify"

var URL_ORDERQUERY = "https://api.mch.weixin.qq.com/pay/orderquery"

var URL_CLOSEORDER = "https://api.mch.weixin.qq.com/pay/closeorder"

var URL_REFUND = "https://api.mch.weixin.qq.com/secapi/pay/refund"

var URL_REFUNDQUERY = "https://api.mch.weixin.qq.com/pay/refundquery"

var URL_DOWNLOADBILL = "https://api.mch.weixin.qq.com/pay/downloadbill"
var URL_MICROPAY = "https://api.mch.weixin.qq.com/pay/micropay"
var URL_SHORTURL = "https://api.mch.weixin.qq.com/tools/shorturl"
var URL_SENDREDPACK = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
var URL_CARD = "https://api.weixin.qq.com/card/create?access_token=%s"
var URL_UPDATE_CARD = "https://api.weixin.qq.com/card/update?access_token=%s"
