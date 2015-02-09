package db

var Table_name_usualmessages_in = "wx_usualmessages_in"
var Wx_usualmessages_in = `CREATE TABLE wx_usualmessages_in (
ToUserName   varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '开发者微信号' ,
FromUserName  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '发送方帐号（一个OpenID）' ,
CreateTime  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息创建时间' ,
MsgType  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息类型' ,
MsgId  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NOT NULL COMMENT '消息id，64位整型' ,
Content  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文本消息内容' ,
PicUrl  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片链接' ,
MediaId  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片消息媒体id，可以调用多媒体文件下载接口拉取数据。' ,
Format  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '语音格式，如amr，speex等' ,
ThumbMediaId  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。' ,
Location_X  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地理位置维度' ,
Location_Y  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地理位置经度' ,
Scale  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地图缩放大小' ,
Label  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地理位置信息' ,
Title  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息标题' ,
Description  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息描述' ,
Url  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息链接' ,
Recognition  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '语音识别结果，UTF8编码' ,
PRIMARY KEY (msgid)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
ROW_FORMAT=COMPACT
;`
var Table_name_pushmessages_in = "wx_pushmessages_in"

var Wx_pushmessages_in = `CREATE TABLE wx_pushmessages_in (
ToUserName  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NOT NULL COMMENT '开发者微信号' ,
FromUserName  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '发送方帐号（一个OpenID）' ,
CreateTime  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息创建时间 （整型）' ,
MsgType  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息类型，event' ,
Event varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '事件类型' ,
EventKey  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '事件KEY值，qrscene_为前缀，后面为二维码的参数值' ,
Ticket  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '二维码的ticket，可用来换取二维码图片' ,
Latitude  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地理位置纬度' ,
Longitude  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地理位置经度' ,
WxPrecision  varchar(255) CHARACTER SET  utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地理位置精度' 
)
;`
