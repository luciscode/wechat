订单关闭 
http://121.40.35.3/closeorder?out_trade_no=tsfa
# 关闭订单#

## 应用场景##

以下情况需要调用关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。

## 接口地址##

#####URL_CLOSEORDER		

	https://api.mch.weixin.qq.com/pay/closeorder

## 请求参数##

	type Closeorderrequest struct {

	XMLName      xml.Name `xml:"xml"`
	Appid        string   `xml:"appid"`                  //公众账号ID
	Mch_id       string   `xml:"mch_id"`                 //商户号
	Nonce_str    string   `xml:"nonce_str"`              //随机字符串
	Out_trade_no string   `xml:"out_trade_no,omitempty"` //商户订单号
	Sign         string   `xml:"sign"`                   //签名
	RequestXML   string   `xml:"-"`                      //最终请求xml串
	}
## 返回参数##

	type Closeorderresponse struct {

	XMLName     xml.Name `xml:"xml"`                  
	Return_code string   `xml:"return_code"`            //返回状态码
	Return_msg  string   `xml:"return_msg,omitempty"`   //返回信息
	//以下字段在return_code为SUCCESS的时候有返回
	Appid        string `xml:"appid"`                   //公众账号ID
	Mch_id       string `xml:"mch_id"`                  //商户号
	Nonce_str    string `xml:"nonce_str"`               //随机字符串
	Sign         string `xml:"sign"`                    //签名
	Err_code     string `xml:"err_code,omitempty"`      //错误代码
	Err_code_des string `xml:"err_code_des,omitempty"`  //错误代码描述

	ReponseXML string `xml:"-"` //结果串

	}
## 开发流程##

1.创建请求结构体

	v Closeorderrequest
 

2.按照需求填充v的数据

	v.Appid=..
 

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到 v.Sign

	v.Signmd5()


4.使用v.Xml()方法生成xml请求串v.RequestXML

    v.Xml()


5.使用v.Dorequest()方法向接口 URL_CLOSEORDER 发送数据v.RequestXML,将得到返回数据解析,得到response Closeorderresponse

    response :=v.Dorequest()




GET
int out_trade_no
out 