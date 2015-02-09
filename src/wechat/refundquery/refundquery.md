http://121.40.35.3/refundquery?out_trade_no=rNPVzm2g1ZQSJMfcH8P3eOrHO4VMbqgR

# 查询退款#

## 应用场景##

提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，用零钱支付的退款20分钟内到账，银行卡支付的退款3个工作日后重新查询退款状态

[微信商户平台]:https://pay.weixin.qq.com/index.php/trade/apply_refund


## 接口地址##
#####URL_REFUNDQUERY
	https://api.mch.weixin.qq.com/pay/refundquery



## 请求参数##

	type Refundqueryrequest struct {

	XMLName        xml.Name `xml:"xml"`
	Appid          string   `xml:"appid"`                    //公众账号ID
	Device_info    string   `xml:"device_info,omitempty"`    //设备号
	Mch_id         string   `xml:"mch_id"`                   //商户号
	Nonce_str      string   `xml:"nonce_str"`                //随机字符串
	Out_refund_no  string   `xml:"out_refund_no,omitempty"`  //商户退款单号
	Out_trade_no   string   `xml:"out_trade_no"`             //商户订单号
	Refund_id      string   `xml:"refund_id,omitempty"`      //微信退款单号
	Sign           string   `xml:"sign"`                     //签名
	Transaction_id string   `xml:"transaction_id,omitempty"` //微信订单号

	RequestXML string `xml:"-"` //最终请求xml串

	}

## 返回参数##

	type Refundqueryresponse struct {

	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`                //返回状态码
	Return_msg  string   `xml:"return_msg"`                 //返回信息
	//以下字段在return_code为SUCCESS的时候有返回
	Result_code  string `xml:"result_code"`                 //业务结果
	Err_code     string `xml:"err_code,omitempty"`          //错误码
	Err_code_des string `xml:"err_code_des,omitempty"`      //错误描述
	Appid        string `xml:"appid"`                       //公众账号ID
	Mch_id       string `xml:"mch_id"`                      //商户号
	Device_info  string `xml:"device_info,omitempty"`       //设备号

	Nonce_str      string `xml:"nonce_str"`                 //随机字符串
	Sign           string `xml:"sign"`                      //签名
	Transaction_id string `xml:"transaction_id"`            //微信订单号
	Out_trade_no   string `xml:"out_trade_no"`              //商户订单号
	Total_fee      string `xml:"total_fee"`                 //订单总金额
	Fee_type       string `xml:"fee_type,omitempty"`        //订单金额货币种类
	Cash_fee       string `xml:"cash_fee"`                  //现金支付金额
	Cash_fee_type  string `xml:"cash_fee_type,omitempty"`   //货币种类

	Refund_fee string `xml:"refund_fee"`                    //退款金额

	Coupon_refund_fee string `xml:"coupon_refund_fee,omitempty"` //代金券或立减优惠退款金额

	Refund_count string `xml:"refund_count"`                //退款笔数
    //后面还有一些字段是关于优惠券的，也是动态变化的，暂时不做处理，改日加上

	ReponseXML string `xml:"-"` //结果串
	}


## 开发流程##

1.创建请求结构体
	
	v Refundqueryrequest
 

2.按照需求填充v的数据

	v.Appid=...
 

3.使用v.Signmd5()方法对v里面的非空字段签名得到v.Sign
	
	v.Signmd5()

4.使用v.Xml()方法生成xml请求串v.RequestXML

	v.Xml()

5.使用v.Dorequest()方法向接口URL_REFUNDQUERY发送数据v.RequestXML,将得到返回数据解析得到response Refundqueryresponse

	reponse：＝v.Dorequest()

