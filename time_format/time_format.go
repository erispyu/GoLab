package main

import (
	"fmt"
	"time"
)

func main() {
	downtimeAnnouncements := "Schedule maintenance from 13/06/23 15:50 to 13/06/23 16:00"
	resumeTimeStr := downtimeAnnouncements[len(downtimeAnnouncements)-14:]
	fmt.Println(resumeTimeStr)
	loc, _ := time.LoadLocation("Asia/Singapore")
	resumeTime, _ := time.ParseInLocation("02/01/06 15:04", resumeTimeStr, loc)
	fmt.Println(resumeTime)
	fmt.Println(resumeTime.Unix())
}
