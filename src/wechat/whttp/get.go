package whttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(geturl string) []byte {

	res, err := http.Get(geturl)
	if err != nil {
		fmt.Println(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)

	}
	return robots
}
