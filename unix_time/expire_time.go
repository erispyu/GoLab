package main

import (
	"fmt"
	"time"
)

const (
	OneSec    = 1
	OneMinute = 60
	OneHour   = 60 * OneMinute
	OneDay    = 24 * OneHour
	OneWeek   = 7 * OneDay
	OneMonth  = 30 * OneDay
)

const (
	ISOFormat      = "2006-01-02T15:04:05Z"
	LocalISOFormat = "2006-01-02T15:04:05.999"
)

func main() {
	// pendingToLocalHour := 24
	// ctime := time.Now().Unix() - 3600*20
	// boundTime := time.Now().Add(-time.Duration(pendingToLocalHour) * OneHour).Unix()
	// println(ctime)
	// println(boundTime)
	// println(ctime < boundTime)
	for {
		s := time.Now().Format(LocalISOFormat)
		fmt.Println(s)
		// random sleep 100 ~ 500 ms
		time.Sleep(time.Duration(100+time.Now().UnixNano()%400) * time.Millisecond)
	}
}
