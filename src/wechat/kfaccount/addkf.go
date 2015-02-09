package kfaccount

import (
	"fmt"
	"net/http"
	"wechat/auth"
	"wechat/whttp"
	"wechat/wjson"
)

var addkfaccount string

func Addkfaccount(w http.ResponseWriter, r *http.Request) {
	a := auth.Access_token(nil, nil)
	addkfaccount = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=" + a.Access_token

	v := Kfaccountrequest{}
	v.Kf_account = "test1@tes"
	v.Nickname = "客服1"
	v.Password = "pswmd5"
	jsonstr, _ := wjson.Encodestruct(v)
	v.RequestJSON = jsonstr
	data := whttp.Post(addkfaccount, v.RequestJSON)
	reponse := Kfaccountresponse{}
	err := wjson.Decodebytes(data, &reponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "结果:"+reponse.Errmsg)

}
