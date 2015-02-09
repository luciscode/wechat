package newspresponse

import (
	"encoding/xml"
	"wechat/wlog"
	"wechat/wtime"
	"wechat/wxml"
)

type Newsitem struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"` //图片地址
	Url         string `xml:"Url"`    //点击打开的网页
}
type Newsresponse struct {
	XMLName      xml.Name   `xml:"xml"`
	ToUserName   string     `xml:"ToUserName"`
	FromUserName string     `xml:"FromUserName"`
	CreateTime   string     `xml:"CreateTime"`
	MsgType      string     `xml:"MsgType"`
	ArticleCount int        `xml:"ArticleCount"`
	NewsItem     []Newsitem `xml:"Articles>item"`

	ReponseXML string `xml:"-"` //结果串

}

func Getnewsresponse(tousername, fromeusername string, newsitem []Newsitem) *Newsresponse {

	v := &Newsresponse{MsgType: "news", CreateTime: wtime.Time_stamp()}
	v.ToUserName = tousername
	v.FromUserName = fromeusername
	v.ArticleCount = len(newsitem)
	for index := 0; index < len(newsitem); index++ {
		v.NewsItem = append(v.NewsItem, newsitem[index])

	}
	v.Xml()
	return v

}
func GetnewsItem(title, description, picurl, url string) Newsitem {
	v := Newsitem{}
	v.Title = title
	v.Description = description
	v.PicUrl = picurl
	v.Url = url
	return v

}
func (v *Newsresponse) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.ReponseXML = xmlresult
	wlog.PrintlnW(xmlresult, true)

	return err

}
