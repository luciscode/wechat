http://121.40.35.3/refund?transaction_id=1000190778201501060009429783&out_trade_no=rNPVzm2g1ZQSJMfcH8P3eOrHO4VMbqgR

# 申请退款#

## 应用场景##

1.交易时间超过半年的订单无法提交退款；

2.微信支付退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。一笔退款失败后重新提交，要采用原来的退款单号。总退款金额不能超过用户实际支付金额。

3.接口提交成功后，还需要在[微信商户平台][]由商户管理员审核退款

[微信商户平台]: https://pay.weixin.qq.com/index.php/trade/apply_refund


## 接口地址##
#####URL_REFUND
	https://api.mch.weixin.qq.com/secapi/pay/refund

## 证书需求##

需要[证书](https://pay.weixin.qq.com/index.php/home/login?return_url=%2Findex.php%2Faccount%2Fapi_cert)

##请求参数##

    type Refundrequest struct {

	XMLName         xml.Name `xml:"xml"`
	Appid           string   `xml:"appid"`                     //公众账号ID
	Device_info     string   `xml:"device_info,omitempty"`     //设备号
	Mch_id          string   `xml:"mch_id"`                    //商户号
	Nonce_str       string   `xml:"nonce_str"`                 //随机字符串
	Op_user_id      string   `xml:"op_user_id"`                //操作员
	Out_refund_no   string   `xml:"out_refund_no"`             //商户退款单号
	Out_trade_no    string   `xml:"out_trade_no"`              //商户订单号
	Refund_fee      string   `xml:"refund_fee"`                //退款金额
	Refund_fee_type string   `xml:"refund_fee_type,omitempty"` //货币种类
	Sign            string   `xml:"sign"`                      //签名
	Total_fee       string   `xml:"total_fee"`                 //总金额
	Transaction_id  string   `xml:"transaction_id"`            //微信订单号

	RequestXML      string   `xml:"-"`                         //最终请求xml串

    }

## 返回参数##

	type Refundreponse struct {

	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`
	Return_msg  string   `xml:"return_msg"`
	//以下字段在return_code为SUCCESS的时候有返回
    Result_code  string `xml:"result_code"`                            //业务结果
	Err_code     string `xml:"err_code,omitempty"`                     //错误代码
	Err_code_des string `xml:"err_code_des,omitempty"`                 //错误代码描述
	Appid                string `xml:"appid"`                          //公众账号ID
	Mch_id               string `xml:"mch_id"`                         //商户号
	Device_info          string `xml:"device_info,omitempty"`          //设备号
	Nonce_str            string `xml:"nonce_str"`                      //随机字符串
	Sign                 string `xml:"sign"`                           //签名
	Transaction_id       string `xml:"transaction_id"`                 //微信订单号
	Out_trade_no         string `xml:"out_trade_no"`                   //商户订单号
	Out_refund_no        string `xml:"out_refund_no"`                  //商户退款单号
	Refund_id            string `xml:"refund_id"`                      //微信退款单号
	Refund_fee           string `xml:"refund_fee"`                     //退款金额
	Refund_fee_type      string `xml:"refund_fee_type,omitempty"`      //退款货币种类
	Total_fee            string `xml:"total_fee"`                      //订单总金额
	Fee_type             string `xml:"fee_type,omitempty"`             //订单金额货币种类
	Cash_fee             string `xml:"cash_fee"`                       //现金支付金额
	Cash_fee_type        string `xml:"cash_fee_type,omitempty"`        //货币种类
	Cash_refund_fee      string `xml:"cash_refund_fee,omitempty"`      //现金退款金额
	Cash_refund_fee_type string `xml:"cash_refund_fee_type,omitempty"` //现金退款货币类型
    //下面被我注释屌的5个字段，不是我不想用，二手里面的$n数值是变化的，还不知道怎么解析变化的xml字段,所以先放着
	// Coupon_refund_fee string `xml:"coupon_refund_fee,omitempty"` //代金券或立减优惠退款金额
	// Coupon_count string `xml:"coupon_count,omitempty"` //代金券或立减优惠使用数量
	// 	coupon_batch_id_$n  string `xml:"coupon_batch_id_$n,omitempty"`            //代金券或立减优惠批次ID
	// coupon_id_$n  string `xml:"coupon_id_$n,omitempty"`            //代金券或立减优惠ID
	// coupon_fee_$n  string `xml:"coupon_fee_$n,omitempty"`            //单个代金券或立减优惠支付金额

	

	ReponseXML string `xml:"-"` //结果xml串

	}


## 开发流程##

1.创建请求结构体

	v Refundrequest
 

2.按照需求填充v的数据

	V.Appid=...
 

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign

	v.Signmd5()


4.使用v.Xml()方法生成xml请求串v.RequestXML

	v.Xml()


5.使用v.DoSSLrequest()方法向接口 URL_REFUND 发送数据v.RequestXML,将得到返回数据解析得到response Refundreponse

 	response :=v.DoSSLrequest()



