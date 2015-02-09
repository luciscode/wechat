# 微信支付请求参数签名工具包#
----------------------
## MD5签名##
	//将要使用的参数存到对应的strcut立马，传入该结构体和API密钥，返回md5签名

	func SignMD5(v interface{}, apikey string) (string, bool) 



