# 网络请求工具包#
----------------------------------------
## Get请求##

	func Get(geturl string) []byte
	

## Post请求##

	func SSLPost(posturl string, data string) []byte





## 带有证书操作的SSLPost请求##
	
    在红包，退款接口用到

	func SSLPost(posturl string, data string) []byte


## 解析请求http.Request

	func Parserequest(r *http.Request) []byte
