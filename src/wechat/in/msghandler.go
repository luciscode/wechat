package in

import (
	"wechat/imagereponse"
	"wechat/textreponse"
	"wechat/videoresponse"
	"wechat/voiceresponse"
)

//处理及时消息
func (v Usualmessage) handleusualmessage() {
	if v.MsgType == "text" {
		//处理关键字
		_, flag := v.Handlekeyword()

		if flag == false { //非关键字处理，原样返回

			response := textreponse.Gettextresponse(v.FromUserName, v.ToUserName, v.Content+"\n测试")

			Writetowx(response.ReponseXML)
		}

	} else if v.MsgType == "image" {
		response := imagereponse.Getimageresponse(v.FromUserName, v.ToUserName, v.MediaId)
		Writetowx(response.ReponseXML)
	} else if v.MsgType == "voice" {
		response := voiceresponse.Getvoiceresponse(v.FromUserName, v.ToUserName, v.MediaId)
		Writetowx(response.ReponseXML)
	} else if v.MsgType == "video" {
		response := videoresponse.Getvideoresponse(v.FromUserName, v.ToUserName, v.MediaId, "test", "testst")
		Writetowx(response.ReponseXML)
	} else if v.MsgType == "music" {
		response := textreponse.Gettextresponse(v.FromUserName, v.ToUserName, "这是一条音乐消息")
		Writetowx(response.ReponseXML)
	}

}

//处理推送消息
func (v Pushmessage) handlepushmessage() {
	if v.Event == "subscribe" {
		response := textreponse.Gettextresponse(v.FromUserName, v.ToUserName, "欢迎关注微信公众号")
		Writetowx(response.ReponseXML)
	}
}
