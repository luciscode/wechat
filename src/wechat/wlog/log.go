package wlog

import (
	"fmt"
	"wechat/config"
)

func Println(msg string) {
	if config.LOG_ON {
		fmt.Println(msg)
	}
}
func PrintlnW(msg string, out bool) {
	if out {
		fmt.Println(msg)
	}
}
