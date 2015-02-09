package in

import (
	"fmt"
	"time"
)

func (v Pushmessage) firstsubcribe() {
	reponsetemplate := "<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%s</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[欢迎关注酷世界公众号]]></Content> </xml>"
	reponse := fmt.Sprintf(reponsetemplate, v.FromUserName, v.ToUserName, time.Now().Unix())

	Writetowx(reponse)
	fmt.Println("对首次关注着进行消息自动发送")

}
