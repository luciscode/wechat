package sendredpack

import (
	"encoding/xml"
	"wechat/config"
	"wechat/whttp"
	"wechat/wlog"
	"wechat/wsign"
	"wechat/wxml"
)

type Sendredpackrequest struct {
	XMLName       xml.Name `xml:"xml"`
	Act_id        string   `xml:"act_id"`                  //maybe
	Act_name      string   `xml:"act_name"`                //maybe
	Client_ip     string   `xml:"client_ip"`               //maybe
	Logo_imgurl   string   `xml:"logo_imgurl,omitempty"`   //maybe
	Max_value     string   `xml:"max_value"`               //maybe
	Mch_billno    string   `xml:"mch_billno"`              //maybe
	Mch_id        string   `xml:"mch_id"`                  //must
	Min_value     string   `xml:"min_value"`               //must
	Nick_name     string   `xml:"nick_name"`               //maybe
	Nonce_str     string   `xml:"nonce_str"`               //must
	Re_openid     string   `xml:"re_openid"`               //must
	Remark        string   `xml:"remark"`                  //maybe
	Send_name     string   `xml:"send_name"`               //must
	Share_content string   `xml:"share_content,omitempty"` //maybe
	Share_imgurl  string   `xml:"share_imgurl,omitempty"`  //maybe
	Share_url     string   `xml:"share_url,omitempty"`     //maybe
	Sign          string   `xml:"sign"`                    //must
	Sub_mch_id    string   `xml:"sub_mch_id,omitempty"`    //must
	Total_amount  string   `xml:"total_amount"`            //maybe
	Total_num     string   `xml:"total_num"`               //maybe
	Wishing       string   `xml:"wishing"`                 //maybe
	Wxappid       string   `xml:"wxappid"`                 //must

	RequestXML string `xml:"-"` //最终请求串

}

func (v *Sendredpackrequest) Xml() error {

	xmlresult, err := wxml.Endoestruct(v)
	v.RequestXML = xmlresult
	wlog.PrintlnW(xmlresult, false)

	return err

}

func (v *Sendredpackrequest) Signmd5() bool {
	sign, _ := wsign.SignMD5(*v, config.API_KEY)
	v.Sign = sign
	wlog.PrintlnW(sign, false)
	return true
}
func (v Sendredpackrequest) DoSSLrequest() Sendredpackreponse {

	data := whttp.SSLPost(config.URL_SENDREDPACK, v.RequestXML)
	sendredpackreponse := Sendredpackreponse{}
	wxml.Decodebytes(data, &sendredpackreponse)
	sendredpackreponse.ReponseXML = string(data)

	return sendredpackreponse

}
