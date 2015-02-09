package wtime

import (
	"strconv"
	"time"
)

func Time_stamp() string {
	v := strconv.FormatInt(time.Now().Unix(), 10)
	return v
}
