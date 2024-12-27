package main

import "time"

const (
	OneSec    = 1
	OneMinute = 60
	OneHour   = 60 * OneMinute
	OneDay    = 24 * OneHour
	OneWeek   = 7 * OneDay
)

func main() {
	expireHours := 3
	currTime := time.Now()
	println(currTime.Unix())
	boundTime := currTime.Add(-time.Duration(expireHours) * time.Hour).Unix()
	println(boundTime)
	ctimeStr := "2023-09-20 10:38:16"
	ctime, _ := time.ParseInLocation("2006-01-02 15:04:05", ctimeStr, time.Local)
	ctimeUnix := ctime.Unix()
	if ctimeUnix < boundTime {
		println("ctime is expired")
	} else {
		println("ctime is not expired")
	}
}
