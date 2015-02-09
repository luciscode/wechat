package whttp

import (
	"io/ioutil"
	"net/http"
)

func Parserequest(r *http.Request) []byte {
	r.ParseForm()
	result, _ := ioutil.ReadAll(r.Body)
	return result
}
