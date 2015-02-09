package test

import (
	"fmt"
	"net/http"
	"wechat/whttp"
)

func Closeordertest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dsa := whttp.Get("http://localhost:80/closeorder?out_trade_no=dsadas")
	fmt.Println(string(dsa))
	fmt.Println("========================================")

}
