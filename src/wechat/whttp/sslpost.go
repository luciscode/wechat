package whttp

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func SSLPost(posturl string, data string) []byte {

	cert, err := tls.LoadX509KeyPair(runtime.GOROOT()+"\\cert.pem", runtime.GOROOT()+"\\key.pem")
	if err != nil {
		fmt.Println(err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	tr := &http.Transport{
		TLSClientConfig: &config,
	}
	client := &http.Client{Transport: tr}

	res, err := client.Post(posturl, "application/x-www-form-urlencoded", strings.NewReader(data))
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
