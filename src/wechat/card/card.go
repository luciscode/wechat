package card

import (
	"fmt"
	"net/http"
	"strconv"
	"wechat/auth"
	"wechat/config"
	"wechat/wtime"
)

func Createcard(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	accesstoken := auth.Access_token(w, r)
	url := fmt.Sprintf(config.URL_CARD, accesstoken.Access_token)

	base_info := Base_infocontent{}
	base_info.Logo_url = "http://www.supadmin.cn/uploads/allimg/120216/1_120216214725_1.jpg"
	base_info.Brand_name = "海底捞"
	base_info.Code_type = "CODE_TYPE_TEXT"
	base_info.Title = "132 元双人火锅套餐"
	base_info.Sub_title = ""
	base_info.Color = "Color081"
	base_info.Notice = "使用时向服务员出示此券"
	base_info.Service_phone = "020-88888888"
	base_info.Description = "不可与其他优惠同享\n如需团购券发票，请在消费时向商户提出\n 店内均可使用，仅限堂食\n 餐前不可打包，餐后未吃完，可打包\n 本团购券不限人数，建议2 人使用，超过建议人数须另收酱料费5 元/位\n 本单谢绝自带酒水饮料"

	date_info := Date_infocontent{}
	date_info.Types = 1
	starttime, _ := strconv.Atoi(wtime.Time_stamp())
	date_info.Begin_timestamp = starttime
	endtime, _ := strconv.Atoi(wtime.Time_stamp())
	date_info.End_timestamp = endtime + 1000000
	base_info.Date_info = date_info

	sku := Skucontent{}
	sku.Quantity = 50000000
	base_info.Sku = sku

	base_info.Use_limit = 1
	base_info.Get_limit = 3
	base_info.Use_custom_code = false
	base_info.Bind_openid = false
	base_info.Can_share = true
	base_info.Url_name_type = "URL_NAME_TYPE_RESERVATION"
	base_info.Custom_url = "http://www.qq.com"
	base_info.Source = "大众点评"

	groupon := Grouponcontent{}
	groupon.Base_info = base_info
	groupon.Deal_detail = "以下锅底2 选1（有菌王锅、麻辣锅、大骨锅、番茄锅、清补凉锅、酸菜鱼锅可选）：\n 大锅1 份12 元\n 小锅2 份16 元\n 以下菜品2 选1\n 特级肥牛1 份30 元\n 洞庭鮰鱼卷1 份20 元\n 其他\n 鲜菇猪肉滑1 份18 元\n 金针菇1 份16 元\n 黑木耳1 份9 元\n 娃娃菜1 份8 元\n 冬瓜1 份6 元\n 火锅面2 个6 元\n 欢乐畅饮2 位12 元\n 自助酱料2 位10 元"
	card := Cardcontent{}
	card.Groupon = groupon
	card.Card_type = "GROUPON"

	request := &Cardrequest{}
	request.Card = card
	request.Json()
	fmt.Println(request.RequestJSON)
	reponse := request.Dorequest(url)

	fmt.Println(reponse.ReponseXML)

}
func Createshakecard(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	accesstoken := auth.Access_token(w, r)
	url := fmt.Sprintf(config.URL_CARD, accesstoken.Access_token)

	base_info := Base_infocontent{}
	base_info.Logo_url = "http://www.supadmin.cn/uploads/allimg/120216/1_120216214725_1.jpg"
	base_info.Brand_name = "微信支付"
	base_info.Code_type = "CODE_TYPE_TEXT"
	base_info.Title = "摇一摇通用券"
	base_info.Sub_title = ""
	base_info.Color = "Color081"
	base_info.Notice = "使用时向服务员出示此券"
	base_info.Service_phone = "020-88888888"
	base_info.Description = "摇一摇，摇出好运来"
	base_info.Get_custom_code_mode = "GET_CUSTOM_CODE_MODE_DEPOSIT"
	base_info.Can_shake = true
	base_info.Shake_slogan_title = "微信支付"
	base_info.Shake_slogan_sub_title = "身体健康，万事如意"

	date_info := Date_infocontent{}
	date_info.Types = 1
	starttime, _ := strconv.Atoi(wtime.Time_stamp())
	date_info.Begin_timestamp = starttime
	endtime, _ := strconv.Atoi(wtime.Time_stamp())
	date_info.End_timestamp = endtime + 1000000
	base_info.Date_info = date_info

	sku := Skucontent{}
	sku.Quantity = 50000000
	base_info.Sku = sku

	base_info.Use_limit = 1
	base_info.Get_limit = 3
	base_info.Use_custom_code = false
	base_info.Bind_openid = false
	base_info.Can_share = true
	base_info.Url_name_type = "URL_NAME_TYPE_RESERVATION"
	base_info.Custom_url = "http://www.qq.com"
	base_info.Source = "微信支付"

	groupon := Grouponcontent{}
	groupon.Base_info = base_info
	groupon.Deal_detail = "送你一个身体健康"
	card := Cardcontent{}
	card.Groupon = groupon
	card.Card_type = "GROUPON"

	request := &Cardrequest{}
	request.Card = card
	request.Json()
	fmt.Println(request.RequestJSON)
	reponse := request.Dorequest(url)

	fmt.Println(reponse.ReponseXML)

}
func Updateshakecard(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	accesstoken := auth.Access_token(w, r)
	url := fmt.Sprintf(config.URL_UPDATE_CARD, accesstoken.Access_token)

	base_info := Base_infoupdatecontent{}

	base_info.Can_shake = true

	cash := Cashcontent{}
	cash.Base_info = base_info

	request := &Updateshakecardrequest{}
	request.Cash = cash
	request.Cardid = "dasdsadas"
	request.Json()
	fmt.Println(request.RequestJSON)
	reponse := request.Dorequest(url)

	fmt.Println(reponse.ReponseXML)

}
