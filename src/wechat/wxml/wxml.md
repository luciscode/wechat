# XML解析工具包#
---
## 解析到struct##
	
	解析xml的byte数据到对应的struct中，在解析whttp请求包里面的http请求返回的时候解析结果用

	func Decodebytes(b []byte, v interface{}) error

## 构造XML串##
 
    根据传入对应的struct，返回要在微信支付请求接口里使用的xml串
	func Endoestruct(v interface{}) (xmlresult string, err error) 

