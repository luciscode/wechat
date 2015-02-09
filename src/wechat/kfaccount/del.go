package kfaccount

import (
	"fmt"
	"net/http"
	"wechat/auth"
	"wechat/whttp"
	"wechat/wjson"
)

var delkfaccount string

func Delkfaccount(w http.ResponseWriter, r *http.Request) {
	a := auth.Access_token(nil, nil)
	delkfaccount = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=" + a.Access_token

	v := Kfaccountrequest{}
	v.Kf_account = "test1@tes"
	v.Nickname = "客服1"
	v.Password = "pswmd5"
	jsonstr, _ := wjson.Encodestruct(v)
	v.RequestJSON = jsonstr
	data := whttp.Post(delkfaccount, v.RequestJSON)
	reponse := Kfaccountresponse{}
	err := wjson.Decodebytes(data, &reponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "结果:"+reponse.Errmsg)

}
