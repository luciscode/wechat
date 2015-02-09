package test

import (
	"fmt"
	"net/http"
	"wechat/whttp"
)

func Micropaytest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dsa := whttp.Get("http://localhost:80/micropay?authcode=dsadas&&fee=1")
	fmt.Println(string(dsa))
	fmt.Println("========================================")

}
