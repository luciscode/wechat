package card

import (
	"wechat/whttp"
	"wechat/wjson"
)

type Skucontent struct {
	Quantity int `json:"quantity"`
}
type Date_infocontent struct {
	Types           int `json:"type"`
	Begin_timestamp int `json:"begin_timestamp"`
	End_timestamp   int `json:"end_timestamp"`
}
type Base_infocontent struct {
	Logo_url         string           `json:"logo_url"`
	Brand_name       string           `json:"brand_name"`
	Code_type        string           `json:"code_type"`
	Title            string           `json:"title"`
	Sub_title        string           `json:"sub_title,omitempty"`
	Color            string           `json:"color"`
	Notice           string           `json:"notice"`
	Service_phone    string           `json:"service_phone,omitempty"`
	Description      string           `json:"description"`
	Date_info        Date_infocontent `json:"date_info"`
	Sku              Skucontent       `json:"sku"`
	Use_limit        int              `json:"use_limit,omitempty"`
	Get_limit        int              `json:"get_limit,omitempty"`
	Use_custom_code  bool             `json:"use_custom_code,omitempty"`
	Bind_openid      bool             `json:"bind_openid,omitempty"`
	Can_share        bool             `json:"can_share,omitempty"`
	Can_give_friend  bool             `json:"can_give_friend,omitempty"`
	Location_id_list string           `json:"location_id_list,omitempty"`
	Url_name_type    string           `json:"url_name_type,omitempty"`
	Custom_url       string           `json:"custom_url,omitempty"`
	Source           string           `json:"source,omitempty"`
	//摇一摇红包
	Get_custom_code_mode   string `json:"get_custom_code_mode,omitempty"`
	Can_shake              bool   `json:"can_shake,omitempty"`
	Shake_slogan_title     string `json:"shake_slogan_title,omitempty"`
	Shake_slogan_sub_title string `json:"shake_slogan_sub_title,omitempty"`
}
type Grouponcontent struct {
	Base_info Base_infocontent `json:"base_info"`

	Deal_detail string `json:"deal_detail"`
}
type Cardcontent struct {
	Card_type string         `json:"card_type"`
	Groupon   Grouponcontent `json:"groupon"`
}

type Cardrequest struct {
	Card        Cardcontent `json:"card"`
	RequestJSON string      `json:"-"`
}

func (v *Cardrequest) Json() error {

	jsonresult, err := wjson.Encodestruct(v)
	v.RequestJSON = jsonresult

	return err

}
func (v Cardrequest) Dorequest(url string) Cardreponse {
	data := whttp.Post(url, v.RequestJSON)

	cardreponse := Cardreponse{}
	wjson.Decodebytes(data, &cardreponse)
	cardreponse.ReponseXML = string(data)

	return cardreponse

}
