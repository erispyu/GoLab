package main

import (
	"fmt"
	"regexp"
)

// const IcOrCompanyIdRegex = `(^[a-zA-Z]\d{9}$)|(^[a-zA-Z]{2}\d{8}$)|(^\d{8}$)`
//
// func main() {
// 	s1 := "R125126386"
// 	icOrCompanyIdRegExp := regexp.MustCompile(IcOrCompanyIdRegex)
// 	fmt.Println(icOrCompanyIdRegExp.MatchString(s1))
// 	s2 := "A 12345678"
// 	fmt.Println(icOrCompanyIdRegExp.MatchString(s2))
// }

func main() {
	// Define the regex pattern
	pattern := `^[\x20-\x7B\x7D-\x7E]*$`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Test strings
	testStrings := []string{
		"A123B",          // Should match
		"Hello | World",  // Should not match
		"Valid String",   // Should match
		"Invalid~String", // Should match
		"·asd",           // Should not match
	}

	// Check each string against the pattern
	for _, str := range testStrings {
		if re.MatchString(str) {
			fmt.Printf("'%s' matches the pattern.\n", str)
		} else {
			fmt.Printf("'%s' does NOT match the pattern.\n", str)
		}
	}
}
