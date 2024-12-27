package main

import "fmt"

func main() {
	var incr int64
	incr = 123456789012
	dateStr := "20210102"
	if incr > 9999999999 {
		panic("Incr value exceeds 10 digits")
	}
	s := fmt.Sprintf("%s%010d", dateStr, incr)
	fmt.Printf("Len: %d, Value: %s\n", len(s), s)
}
