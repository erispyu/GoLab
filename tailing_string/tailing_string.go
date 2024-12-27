package main

import "fmt"

func main() {
	s1 := "12345678"
	s2 := "12345"
	fmt.Println(s1[len(s1)-6:])
	fmt.Println(s2[len(s2)-6:])
}
