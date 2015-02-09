package in

import (
	"fmt"
	"wechat/mgo"
	"wechat/newspresponse"
	"wechat/textreponse"
)

func (v Usualmessage) Handlekeyword() (string, bool) {
	key, err := mgo.MongoFindKeywordLocal(v.Content)
	if err == nil {
		response := textreponse.Gettextresponse(v.FromUserName, v.ToUserName, key.Content)

		Writetowx(response.ReponseXML)
		fmt.Println("对关键字[" + v.Content + "]进行了固定的回复")

		return response.ReponseXML, true
	} else if v.Content == "news" {
		newsitem1 := newspresponse.GetnewsItem("test", "testdescription", "", "http://www.qq.com")
		newsitem2 := newspresponse.GetnewsItem("test2", "testdescription2", "", "")
		var newsitem []newspresponse.Newsitem
		newsitem = append(newsitem, newsitem1)
		newsitem = append(newsitem, newsitem2)

		response := newspresponse.Getnewsresponse(v.FromUserName, v.ToUserName, newsitem)
		Writetowx(response.ReponseXML)
		return response.ReponseXML, true

	} else {
		return "", false
	}

}
