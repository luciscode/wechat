title: 被扫支付
date: 2015-01-20 22:57:30
tags: 被扫支付
---

[源码地址](https://gitcafe.com/lucis/wechat/blob/master/src/wechat/micropay/micropay.md)

## 应用场景 ##

收银员使用扫码设备读取微信用户刷卡授权码以后，二维码或条码信息传送至商户收银台，由商户收银台或者商户后台调用该接口发起支付。当然也可以在收银台发起，但不建议这样做，网络，安全均无法保证
-	由商家发起扣款
-	支付金额小于等于300，且当日支付笔数小于10笔，不需要用户输入密码

<!--more-->
## 接口地址##

#####URL_MICROPAY

	https://api.mch.weixin.qq.com/pay/micropay
## 请求方式
	POST方式
## 请求参数## 

`带有omitempty属性表示该字段属于选填参数，不带omitempty属性表示该字段是必填参数`


	type Mircopayrequest struct {
	XMLName          xml.Name `xml:"xml"`
	Appid            string   `xml:"appid"`                 //公众账号ID
	Attach           string   `xml:"attach,omitempty"`      //附加数据
	Auth_code        string   `xml:"auth_code"`             //授权码
	Body             string   `xml:"body"`                  //商品描述
	Detail           string   `xml:"detail,omitempty"`      //商品详情
	Device_info      string   `xml:"device_info,omitempty"` //设备号
	Fee_type         string   `xml:"fee_type,omitempty"`    //货币类型
	Goods_tag        string   `xml:"goods_tag,omitempty"`   //商品标记
	Mch_id           string   `xml:"mch_id"`                //商户号
	Nonce_str        string   `xml:"nonce_str"`             //随机字符串
	Out_trade_no     string   `xml:"out_trade_no"`          //商户订单号
	Sign             string   `xml:"sign"`                  //签名
	Spbill_create_ip string   `xml:"spbill_create_ip"`      //终端IP
	Time_expite      string   `xml:"time_expite,omitempty"` //交易失效时间
	Time_start       string   `xml:"time_start,omitempty"`  //交易起始时间
	Total_fee        string   `xml:"total_fee"`             //总金额

	RequestXML string `xml:"-"` //最终请求串

	}

## 返回参数##
	type Mircopayresponse struct {
	XMLName        xml.Name `xml:"xml"`
	Return_code    string   `xml:"return_code"`           //返回状态码
	Return_msg     string   `xml:"return_msg"` 		      //返回信息
	Appid          string   `xml:"appid"`                 //公众账号ID
	Mch_id         string   `xml:"mch_id"`                //商户号
	Device_info    string   `xml:"device_info,omitempty"` //设备号
	Nonce_str      string   `xml:"nonce_str"`             //随机字符串
	Sign           string   `xml:"sign"`                  //签名
	Result_code    string   `xml:"result_code"`           //业务结果
	Err_code       string   `xml:"err_code"`              //错误代码
	Err_code_des   string   `xml:"err_code_des"`          //错误代码描述
	Is_subscribe   string   `xml:"is_subscribe"`   		  //是否关注公众账号
	Trade_type     string   `xml:"trade_type"`			  //交易类型
	Bank_type      string   `xml:"bank_type"`			  //付款银行
	Fee_type       string   `xml:"fee_type"`		      //货币类型
	Total_fee      string   `xml:"total_fee"`             //总金额
	Cash_fee_type  string   `xml:"cash_fee_type"`		  //现金支付货币类型
	Cash_fee       string   `xml:"cash_fee"`			  //现金支付金额
	Coupon_fee     string   `xml:"coupon_fee"`			  //代金券或立减优惠金额
	Transaction_id string   `xml:"transaction_id"`		  //微信支付订单号
	Out_trade_no   string   `xml:"out_trade_no"`          //商户订单号
	Attach         string   `xml:"attach"`                //商家数据包
	Time_end       string   `xml:"time_end"`			  //支付完成时间

	ReponseXML string `xml:"-"` //结果xml串

	}
## 开发流程##

1.创建请求结构体

	v Mircopayrequest
 

2.按照需求填充v的数据

	v.Appid=...

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign，签名方式为MD5。

	v.Signmd5()


4.使用v.Xml()生成xml请求串v.RequestXML

	v.Xml()


5.使用v.Dorequest()向接口 URL_MICROPAY 发送数据v.RequestXML,将得到返回数据解析得到response Mircopayresponse

	response :=v.Dorequest()


## 支付结果判断##


    reponse.Return_code为SUCCESS，response.Result_code为SUCCESS,表示支付成功 
   
    reponse.Return_code为SUCCESS，reponse.Result_code为FAIL,reponse.Err_code为USERPAYING 表示正在输入密码

## 更多reponse.Err_code##


| 名称      		|  			  描述 | 		原因  |解决方案 |
| :-------- | --------:| :--: |
| SYSTEMERROR	  	| 	   接口返回错误|  系统超时    | 请立即调用被扫订单结果查询API，查询当前订单状态，并根据订单的状态决定下一步的操作|
| ORDERPAID	  	| 	   订单已支付|  订单号重复    | 请确认该订单号是否重复支付，如果是新单，请使用新订单号提交|
| NOAUTH	  	| 	   商户无权限|  商户没有开通被扫支付权限    | 请开通商户号权限。请联系产品或商务申请|
| AUTHCODEEXPIRE	  	| 	   二维码已过期，请用户在微信上刷新后再试|  用户的条码已经过期    | 	请收银员提示用户，请用户在微信上刷新条码，然后请收银员重新扫码。 直接将错误展示给收银员
|
| NOTENOUGH	  	| 	   余额不足|  用户的零钱余额不足    |请收银员提示用户更换当前支付的卡，然后请收银员重新扫码。建议：商户系统返回给收银台的提示为“用户余额不足.提示用户换卡支付”|
| NOTSUPORTCARD	  	| 	   不支持卡类型|  用户使用卡种不支持当前支付形式,请用户重新选择卡种    | 建议：商户系统返回给收银台的提示为“该卡不支持当前支付，提示用户换卡支付或绑新卡支付”|
| ORDERCLOSED	  	| 	   订单已关闭|  该订单已关    | 商户订单号异常，请重新下单支付|
| ORDERREVERSED	  	| 	   订单已撤销|  当前订单已经被撤销    | 当前订单状态为“订单已撤销”，请提示用户重新支付|
| BANKERROR	  	| 	   银行系统异常|  银行端超时    | 请立即调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作|
| USERPAYING	  	| 	   用户支付中，需要输入密码|  该笔交易因为业务规则要求，需要用户输入支付密码    | 等待5秒，然后调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作|
| AUTH_CODE_ERROR	  	| 	   授权码参数错误|  请求参数未按指引进行填写    | 每个二维码仅限使用一次，请刷新再试|
| AUTH_CODE_INVALID	  	| 	   授权码检验错误|  收银员扫描的不是微信支付的条码    | 请扫描微信支付被扫条码/二维码|
| XML_FORMAT_ERROR	  	| 	   XML格式错误|  XML格式错误    | 请检查XML参数格式是否正确|
| REQUIRE_POST_METHOD	  	| 	   请使用post方法|  未使用post传递参数    | 请检查请求参数是否通过post方法提交|
| SIGNERROR	  	| 	   签名错误|  参数签名结果不正确    | 请检查签名参数和方法是否都符合签名算法要求|
| LACK_PARAMS	  	| 	   缺少参数|  缺少必要的请求参数    | 请检查参数是否齐全|
| NOT_UTF8	  	| 	   编码格式错误|  未使用指定编码格式    | 请使用NOT_UTF8编码格式|
| BUYER_MISMATCH	  	| 	   支付帐号错误|  暂不支持同一笔订单更换支付方    | 请确认支付方是否相同|
| APPID_NOT_EXIST	  	| 	   APPID不存在|  参数中缺少APPID    | 请检查APPID是否正确|
| MCHID_NOT_EXIST	  	| 	   MCHID不存在|  参数中缺少MCHID    | 请检查MCHID是否正确|
| OUT_TRADE_NO_USED	  	| 	   商户订单号重复|  同一笔交易不能多次提交    | 请核实商户订单号是否重复提交|
| APPID_MCHID_NOT_MATCH	  	| 	   appid和mch_id不匹配|  appid和mch_id不匹配    | 请确认appid和mch_id是否匹配|


## [签名校验工具](http://mch.weixin.qq.com/wiki/tools/signverify/)
