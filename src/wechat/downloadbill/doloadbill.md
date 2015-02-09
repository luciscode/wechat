http://121.40.35.3/downloadbill
# 下载对账单#

## 应用场景##

商户可以通过该接口下载历史交易清单。比如掉单、系统错误等导致商户侧和微信侧数据不一致，通过对账单核对后可校正支付状态


1.微信侧未成功下单的交易不会出现在对账单中。支付成功后撤销的交易会出现在对账单中，跟原支付单订单号一致，bill_type为REVOKED

2.微信在次日9点启动生成前一天的对账单，建议商户10点后再获取

3.对账单中涉及金额的字段单位为“元”


## 接口地址##
#####URL_DOWNLOADBILL

	https://api.mch.weixin.qq.com/pay/downloadbill

## 请求参数##

	type Downloadbillrequest struct {

	XMLName     xml.Name `xml:"xml"`
	Appid       string   `xml:"appid"`                 //公众账号ID
	Bill_date   string   `xml:"bill_date"`             //对账单日起
	Bill_type   string   `xml:"bill_type,omitempty"`   //账单类型
	Device_info string   `xml:"device_info,omitempty"` //设备号
	Mch_id      string   `xml:"mch_id"`                //商户号
	Nonce_str   string   `xml:"nonce_str"`             //随机字符串
	Sign        string   `xml:"sign"`                  //签名
	RequestXML  string   `xml:"-"`                     //最终请求串

	}
## 返回参数##

	type Downloadbillresponse struct {

	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`         //返回状态码
	Return_msg  string   `xml:"return_msg"`          //返回信息

	ReponseXML string `xml:"-"` //下载失败时候的结果xml串
	ReponseSTR string `xml:"-"` //下载成功时候的账单字符串
	}


## 开发流程##

1.创建请求结构体

	v Downloadbillrequest
 

2.按照需求填充v的数据
 	
	v.Appid=...

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign

	v.Signmd5()	

4.使用v.Xml()方法生成xml请求串v.RequestXML

	v.Xml()

5.使用v.Dorequest()方法向接口 URL_DOWNLOADBILL 发送数据v.RequestXML,将得到返回数据解析得到response Downloadbillresponse

	reponse :=v.Dorequest()