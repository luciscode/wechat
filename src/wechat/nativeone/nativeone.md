模式1
http://121.40.35.3/native1
# 查询订单#

## 应用场景##
该接口提供所有微信支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑

需要调用查询接口的情况：

当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知

调用支付接口后，返回系统错误或未知交易状态情况

调用被扫支付API，返回USERPAYING的状态

调用关单接口API之前，需确认支付状态

## 接口地址##
https://api.mch.weixin.qq.com/pay/orderquery

## 请求参数##
## 返回参数##
## 开发流程##

1.创建请求结构体:v Refundrequest
 

2.按照需求填充v的数据
 

3.使用v.Signmd5()方法对v里面的非空字段进行签名得到v.Sign


4.使用v.Xml()方法生成xml请求串v.RequestXML


5.使用v.DoSSLrequest()方法向接口发送数据v.RequestXML,将得到返回数据使用wxml.Decodebytes()解析, 得到res Refundreponse