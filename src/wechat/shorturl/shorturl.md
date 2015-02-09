http://121.40.35.3/shorturl
# 转换短链接#
## 应用场景##
该接口主要用于扫码原生支付模式一中的二维码链接转成短链接(weixin://wxpay/s/XXXXXX)，减小二维码数据量，提升扫描速度和精确度。
## 接口地址##

#####URL_SHORTURL

	https://api.mch.weixin.qq.com/tools/shorturl

## 请求参数##

	type Shorturlrequest struct {

	XMLName    xml.Name `xml:"xml"`
	Appid      string   `xml:"appid"`     //公众账号ID
	Long_url   string   `xml:"long_url"`  //URL链接
	Mch_id     string   `xml:"mch_id"`    //商户号
	Nonce_str  string   `xml:"nonce_str"` //随机字符串
	Sign       string   `xml:"sign"`      //签名

	RequestXML string   `xml:"-"`         //最终请求xml串

	}

## 返回结果##

	type Shorturlresponse struct {

	XMLName     xml.Name `xml:"xml"`
	Return_code string   `xml:"return_code"`      //返回状态码
	Return_msg  string   `xml:"return_msg"`       //返回信息
	//以下字段在return_code为SUCCESS的时候有返回

	Appid  string `xml:"appid"`                   //公众账号ID
	Mch_id string `xml:"mch_id"`                  //商户号
	Nonce_str   string `xml:"nonce_str"`          //随机字符串
	Sign        string `xml:"sign"`               //签名
	Result_code string `xml:"result_code"`        //业务结果
	Err_code    string `xml:"err_code,omitempty"` //错误代码
	Short_url   string `xml:"short_url"`          //URL链接

	ReponseXML string `xml:"-"` //结果xml串
	}

## 错误码##
		SIGNERROR	签名错误		参数签名结果不正确		请检查签名参数和方法是否都符合签名算法要求

		REQUIRE_POST_METHOD		请使用post方法		未使用post传递参数 		请检查请求参数是否通过post方法提交

		APPID_NOT_EXIST		APPID不存在		参数中缺少APPID		请检查APPID是否正确

		MCHID_NOT_EXIST		MCHID不存在		参数中缺少MCHID		请检查MCHID是否正确

		APPID_MCHID_NOT_MATCH		appid和mch_id不匹配		appid和mch_id不匹配		请确认appid和mch_id是否匹配

		LACK_PARAMS		缺少参数		缺少必要的请求参		请检查参数是否齐全

		XML_FORMAT_ERROR		XML格式错误		XML格式错误		请检查XML参数格式是否正确

		POST_DATA_EMPTY		post数据为空		post数据不能为空		请检查post数据是否为空

## 开发流程##

1.创建请求结构体
	
	v Shorturlrequest
 

2.按照需求填充v的数据
 	
	v.Appid=...


3.使用v.Signmd5()对v里面的非空字段进行签名得到v.Sign

	v.Signmd5()

4.使用v.Xml()生成xml请求串v.RequestXML
	
	v.Xml()

5.使用v.Dorequest()方法向接口 URL_SHORTURL 发送数据v.RequestXML,将得到返回数据解析得到response Shorturlresponse

 	response :=v.DoSSLrequest()
