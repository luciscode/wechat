#JSAPI支付常见问题


- 支付域名未授权，无法发起该笔交易：支付目录在多个公众号内重复设置，检查一下


- invalid appid：未设置白名单，或支付目录不对，或不是在该公众号内发起，特别注意微信支付H5页面只能在该支付公众号内打开

- JSAPI支付必须设置支付目录，且正式的支付目录必须通过备案，设置规则：如设置支付目录为http://www.wechatpay.com/,那么使用JSAPI调起微信支付的页面必须是 http://www.wechatpay.com/xxx，而不能是http://www.wechatpay.com/xxx/xxx

- JSAPI支付可以设置测试目录，设置规则同上，但这里不限制域名类型，可以是未备案的域名也可是IP地址，同时必须把要发起测试支付的微信号加入白名单


# JSAPI支付#


## 应用场景##

用户通过消息或扫描二维码在微信内打开网页时，可以调用微信支付完成下单购买的流程

## 接口##

仅可在微信内部浏览器内使用的JS接口：getBrandWCPayRequest

## 请求参数##

	type Jsapirequest struct {

	XMLName   xml.Name `xml:"xml"`
	Appid     string   `xml:"appid"`     //公众账号ID
	TimeStamp string   `xml:"timeStamp"` //对账单日起
	Package   string   `xml:"package,"`  //账单类型
	SignType  string   `xml:"signType"`  //商户号
	Nonce_str string   `xml:"nonce_str"` //随机字符串
	PaySign   string   `xml:"paySign"`   //签名

	}

## 返回参数##


err_msg	get_brand_wcpay_request:ok	支付成功
get_brand_wcpay_request:cancel	支付过程中用户取消
get_brand_wcpay_request:fail	支付失败


## 开发流程##

1.创建请求结构体:v Unifieldrequest
 

2.按照需求填充v的数据
 

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign

	v.Signmd5()


4.使用v.Xml()方法生成xml请求串v.RequestXML
	
	v.Xml()
  	

5.使用v.Dorequest()方法向统一下单接口URL_UNIFIEDORDER发送数据 v.RequestXML,将得到返回数据解析得到response 
Unifiedorderreponse,里面包含Prepay_id
	
	response :=v.Dorequest
	
6.创建jv Jsapirequest,填充数据

	jv.package为prepay_id=res.Prepay_id


7.使用jv.Signmd5()方法对v里面的非空字段进行签名得到v.PaySign

	jv.Signmd5()

9.jv里面的参数传入JS函数 getBrandWCPayRequest，调起微信支付，完成支付


示例代码如下：

	function onBridgeReady(){

	WeixinJSBridge.invoke(

       'getBrandWCPayRequest', {
           "appId" : "wx2421b1c4370ec43b",     //公众号名称，由商户传入     
           "timeStamp":" 1395712654",         //时间戳，自1970年以来的秒数     
           "nonceStr" : "e61463f8efa94090b1f366cccfbbb444", //随机串     
           "package" : "prepay_id=u802345jgfjsdfgsdg888",     
           "signType" : "MD5",         //微信签名方式:     
           "paySign" : "70EA570631E4BB79628FBCA90534C63FF7FADD89" //微信签名 
       },

       function(res){     
           if(res.err_msg == "get_brand_wcpay_request:ok" ) {}     // 使用以上方式判断前端返回,微信团队郑重提示：res.err_msg将在用户支付成功后返回    ok，但并不保证它绝对可靠。 
       }
	); 
	}

	if (typeof WeixinJSBridge == "undefined"){

		if( document.addEventListener ){

        document.addEventListener('WeixinJSBridgeReady', onBridgeReady, false);
		
		   }else if (document.attachEvent){

        document.attachEvent('WeixinJSBridgeReady', onBridgeReady); 
        document.attachEvent('onWeixinJSBridgeReady', onBridgeReady);

	}
	}else{
	    onBridgeReady();

	}
