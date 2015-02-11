package whttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wechat/wlog"
)

func Post(posturl string, data string) []byte {

	res, err := http.Post(posturl, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)

	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)

	}
	wlog.PrintlnW(string(robots), false)
	return robots

}
