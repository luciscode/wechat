package card

import (
	"wechat/whttp"
	"wechat/wjson"
)

type Base_infoupdatecontent struct {
	Can_shake bool `json:"can_shake,omitempty"`
}
type Cashcontent struct {
	Base_info Base_infoupdatecontent `json:"base_info"`
}

type Updateshakecardrequest struct {
	Cardid      string      `json:"card_id"`
	Cash        Cashcontent `json:"cash"`
	RequestJSON string      `json:"-"`
}

func (v *Updateshakecardrequest) Json() error {

	jsonresult, err := wjson.Encodestruct(v)
	v.RequestJSON = jsonresult

	return err

}
func (v Updateshakecardrequest) Dorequest(url string) Updateshakecardresponse {
	data := whttp.Post(url, v.RequestJSON)

	updateshakecardresponse := Updateshakecardresponse{}
	wjson.Decodebytes(data, &updateshakecardresponse)
	updateshakecardresponse.ReponseXML = string(data)

	return updateshakecardresponse

}
