package test

import (
	"fmt"
	"net/http"
	"wechat/whttp"
)

func Testpost(w http.ResponseWriter, r *http.Request) {

	url := "http://www.sanjiang.com/test.aspx"
	// gettest := whttp.Get(url + "?gettest=wechatget")
	// fmt.Println(string(gettest))
	posttest := whttp.Post(url, "posttest=wechatpost")

	fmt.Println(string(posttest))

}
