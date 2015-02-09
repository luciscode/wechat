title: 统一下单
date: 2015-01-21 11:34:55
tags:
---
[源码地址](https://gitcafe.com/lucis/wechat/blob/master/src/wechat/unifiedorder/unifiedorder.md)

## 应用场景##

-	除被扫支付场景以外，其余的支付方式都要先调用该接口获取预支付prepay_id

-	扫码、JSAPI、APP等不同场景使用获取的prepay_id生成交易串调起微信支付。

<!--more-->

## 接口链接##

#####URL_UNIFIEDORDER

	https://api.mch.weixin.qq.com/pay/unifiedorder

## 请求方式
	POST方式

## 请求参数##

`带有omitempty属性表示该字段属于选填参数，不带omitempty属性表示该字段是必填参数`


	type Unifieldrequest struct {

	XMLName          xml.Name `xml:"xml"`                   
	Appid            string   `xml:"appid"`                 //公众账号ID
	Attach           string   `xml:"attach,omitempty"`      //附加数据
	Body             string   `xml:"body"`                  //商品描述
	Detail           string   `xml:"detail,omitempty"`      //商品详情
	Device_info      string   `xml:"device_info,omitempty"` //设备号
	Fee_type         string   `xml:"fee_type,omitempty"`    //货币类型
	Notify_url       string   `xml:"notify_url"`            //通知地址
	Goods_tag        string   `xml:"goods_tag,omitempty"`   //商品标记
	Mch_id           string   `xml:"mch_id"`                //商户号
	Nonce_str        string   `xml:"nonce_str"`             //随机字符串
	Openid           string   `xml:"openid,omitempty"`      //用户标识
	Out_trade_no     string   `xml:"out_trade_no"`          //商户订单号
	Product_id       string   `xml:"product_id,omitempty"`  //商品ID
	Sign             string   `xml:"sign"`                  //签名
	Spbill_create_ip string   `xml:"spbill_create_ip"`      //终端IP
	Time_expire      string   `xml:"time_expire,omitempty"` //交易结束时间
	Time_start       string   `xml:"time_start,omitempty"`  //交易起始时间
	Total_fee        string   `xml:"total_fee"`             //总金额
	Trade_type       string   `xml:"trade_type"`            //交易类型

	RequestXML string `xml:"-"` //最终请求xml串

	}

## 返回参数##

	type Unifiedorderreponse struct {

	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`          //返回状态码
	Return_msg  string   `xml:"return_msg"`           //返回信息
    //以下字段在return_code为SUCCESS的时候有返回
	Appid        string `xml:"appid"`                 //公众账号ID
	Mch_id       string `xml:"mch_id"`                //商户号
	Device_info  string `xml:"device_info,omitempty"` //设备号
	Nonce_str    string `xml:"nonce_str"`             //随机字符串
	Sign         string `xml:"sign"`                  //签名
	Result_code  string `xml:"result_code"`           //业务结果
	Err_code     string `xml:"err_code,omitempty"`    //错误代码
	Err_code_des string `xml:"err_code_des,omitempty"`//错误代码描述
    //以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Trade_type string `xml:"trade_type,omitempty"`    //交易类型
	Prepay_id  string `xml:"prepay_id,omitempty"`     //预支付交易会话标识
	Code_url   string `xml:"code_url,omitempty"`      //二维码链接

	ReponseXML string `xml:"-"` //结果xml串

	}

## 开发流程##


1.创建请求结构体

	v Unifieldrequest
 

2.按照需求填充v的数据
 
	v.Appid=...

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign

	v.Signmd5()


4.使用v.Xml()生成xml请求串v.RequestXML

	v.Xml()


5.使用v.Dorequest()向接口config.URL_UNIFIEDORDER发送数据v.RequestXML,将得到返回数据解析得到response Unifiedorderreponse

	response :=v.Dorequest()

## 下单结果判断##


    reponse.Return_code为SUCCESS，response.Result_code为SUCCESS,表示获取预支付订单成功
   
## 更多reponse.Err_code##
| 名称      |描述	    |原因		|解决方案|
| :-------- | --------:| ---: |:---: |
|NOAUTH		|商户无此接口权限	|商户未开通此接口权限|请商户前往申请此接口权限|
|NOTENOUGH		|余额不足|用户帐号余额不足|用户帐号余额不足，请用户充值或更换支付卡后再支付|
|ORDERPAID		|商户订单已支付	|商户订单已支付，无需重复操作|商户订单已支付，无需更多操作|
|ORDERCLOSED		|订单已关闭	|当前订单已关闭，无法支付|当前订单已关闭，请重新下单|
|SYSTEMERROR		|系统错误	|系统超时|系统异常，请用相同参数重新调用|
|APPID_NOT_EXIST		|APPID不存在	|参数中缺少APPID|请检查APPID是否正确|
|MCHID_NOT_EXIST		|MCHID不存在	|参数中缺少MCHID|请检查MCHID是否正确|
|APPID_MCHID_NOT_MATCH		|appid和mch_id不匹配	|appid和mch_id不匹配|请确认appid和mch_id是否匹配|
|LACK_PARAMS		|缺少参数	|缺少必要的请求参数|请检查参数是否齐全|
|OUT_TRADE_NO_USED		|商户订单号重复	|同一笔交易不能多次提交|请核实商户订单号是否重复提交|
|SIGNERROR		|签名错误	|参数签名结果不正确|请检查签名参数和方法是否都符合签名算法要求|
|XML_FORMAT_ERROR		|XML格式错误	|XML格式错误|请检查XML参数格式是否正确|
|REQUIRE_POST_METHOD		|请使用post方法	|未使用post传递参数 |请检查请求参数是否通过post方法提交|
|POST_DATA_EMPTY		|post数据为空	|post数据不能为空|请检查post数据是否为空|
|NOT_UTF8		|编码格式错误	|未使用指定编码格式|请使用NOT_UTF8编码格式|


## [签名校验工具](http://mch.weixin.qq.com/wiki/tools/signverify/)
