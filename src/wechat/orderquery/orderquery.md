查询:http://121.40.35.3/orderquery?out_trade_no=rNPVzm2g1ZQSJMfcH8P3eOrHO4VMbqgR

# 查询订单#

## 应用场景##
该接口提供所有微信支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑

需要调用查询接口的情况：

当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知

调用支付接口后，返回系统错误或未知交易状态情况

调用被扫支付API，返回USERPAYING的状态

调用关单接口API之前，需确认支付状态

## 接口地址

#####URL_ORDERQUERY

	https://api.mch.weixin.qq.com/pay/orderquery

## 请求参数##

type Orderqueryrequest struct {

	XMLName        xml.Name `xml:"xml"`
	Appid          string   `xml:"appid"`                    //公众账号ID
	Mch_id         string   `xml:"mch_id"`                   //商户号
	Nonce_str      string   `xml:"nonce_str"`                //随机字符串
	Out_trade_no   string   `xml:"out_trade_no,omitempty"`   //商户订单号
	Sign           string   `xml:"sign"`                     //签名
	Transaction_id string   `xml:"transaction_id,omitempty"` //微信订单号
	RequestXML     string   `xml:"-"`                        //最终请求串

}

## 返回参数##

	type Orderqueryresponse struct {

	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`                 //返回状态码
	Return_msg  string   `xml:"return_msg"`                  //返回信息
	//以下字段在return_code为SUCCESS的时候有返回
	Appid        string `xml:"appid"`                        //公众账号ID
	Mch_id       string `xml:"mch_id"`                       //商户号
	Nonce_str    string `xml:"nonce_str"`                    //随机字符串
	Sign         string `xml:"sign"`                         //签名
	Result_code  string `xml:"result_code"`                  //业务结果
	Err_code     string `xml:"err_code,omitempty"`           //错误代码
	Err_code_des string `xml:"err_code_des,omitempty"`       //错误代码描述

	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Device_info  string `xml:"device_info,omitempty"`        //设备号
	Openid       string `xml:"openid"`                       //用户标识
	Is_subscribe string `xml:"is_subscribe"`                 //是否关注公众账号
	Trade_type   string `xml:"trade_type"`                   //交易类型
	Trade_state  string `xml:"trade_state"`                  //交易状态

	Bank_type     string `xml:"bank_type"`                   //付款银行
	Total_fee     string `xml:"total_fee"`                   //总金额
	Fee_type      string `xml:"fee_type,omitempty"`          //货币种类
	Cash_fee      string `xml:"cash_fee"`                    //现金支付金额
	Cash_fee_type string `xml:"cash_fee_type,omitempty"`     //现金支付货币类型
	// Coupon_fee   string `xml:"coupon_fee,omitempty"`      //代金券或立减优惠金额
	// Coupon_count   string `xml:"coupon_count,omitempty"`  //代金券或立减优惠使用数量
	// Coupon_batch_id_$n   string `xml:"coupon_batch_id_$n,omitempty"`           //代金券或立减优惠批次ID
	// Coupon_id_$n   string `xml:"coupon_id_$n,omitempty"`           //代金券或立减优惠ID
	// Coupon_fee_$n   string `xml:"coupon_fee_$n,omitempty"`           //单个代金券或立减优惠支付金额
	Transaction_id string `xml:"transaction_id"`             //微信支付订单号
	Out_trade_no   string `xml:"out_trade_no"`               //商户订单号
	Attach         string `xml:"attach,omitempty"`           //商家数据包
	Time_end       string `xml:"time_end"`                   //支付完成时间

	ReponseXML string `xml:"-"` //结果串
	}
## 开发流程##


1.创建请求结构体

	v Orderqueryrequest
 

2.按照需求填充v的数据
 
	v.Appid=...

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign

	v.Signmd5()


4.使用v.Xml()生成xml请求串v.RequestXML

	v.Xml()


5.使用v.Dorequest()向接口 URL_ORDERQUERY 发送数据v.RequestXML,将得到返回数据解析得到response Orderqueryresponse

	response :=v.Dorequest()