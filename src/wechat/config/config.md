#配置文件
##微信支付信息
公众号的为唯一标识，在[公众号平台](mp.weixin.qq.com) 里面的开发者中心查看

	var APP_ID = "wx*****************" 



公众号的密钥，和APPID对应，也在[公众号平台](mp.weixin.qq.com) 里面的开发者中心查看

	var APP_SECRET = "**********************************"


微信支付商户号，在微信支申请通过够，会由微信邮件发送到申请微信支付时候填写的邮箱，如果确定申请通过，但没看到邮件的话，在自己的邮箱里面搜关键字:weixin ，可能被丢到垃圾箱了

	var MCH_ID = "10******"   


微信支付子商户号，受理模式使用

	var SUB_MCH_ID = ""



API密钥，所有的支付接口签名时候到要使用，该密钥需要去设置，使用的时候必须保证是设置的该API密钥，密钥为不大于32位的随机字符串，自己生成即可

	var API_KEY = "**********************************"

##支付接口

统一下单接口


	var URL_UNIFIEDORDER = "https://api.mch.weixin.qq.com/pay/unifiedorder"



支付通知接口，异步接受微信支付的支付结果

	var URL_UNIFIEDORDER_NOTIFYURL = "http://121.40.35.3/notify"

查询订单接口

	var URL_ORDERQUERY = "https://api.mch.weixin.qq.com/pay/orderquery"

关闭订单接口

	var URL_CLOSEORDER = "https://api.mch.weixin.qq.com/pay/closeorder"

退款接口

	var URL_REFUND = "https://api.mch.weixin.qq.com/secapi/pay/refund"

查询退款接口

	var URL_REFUNDQUERY = "https://api.mch.weixin.qq.com/pay/refundquery"

下载对账单接口

	var URL_DOWNLOADBILL = "https://api.mch.weixin.qq.com/pay/downloadbill"

被扫支付接口

	var URL_MICROPAY = "https://api.mch.weixin.qq.com/pay/micropay"

转换短链接接口

	var URL_SHORTURL = "https://api.mch.weixin.qq.com/tools/shorturl"

发送红包接口

	var URL_SENDREDPACK = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
