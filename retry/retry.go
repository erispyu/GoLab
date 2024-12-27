package main

import (
	"fmt"
	"time"
)

func main() {
	maxRetries := 3
	retryIntervals := []int{1, 2, 4}
	currRetries := 0
	for currRetries <= maxRetries {
		fmt.Println(time.Now().String())
		if currRetries == maxRetries {
			break
		}
		time.Sleep(time.Duration(retryIntervals[currRetries]) * time.Second)
		currRetries++
	}

}
