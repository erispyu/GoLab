package main

import "fmt"

func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

func main() {
	// logHeader := "[CheckBiFastPayoutAccountInquiry]"
	param := map[string]string{}
	param["a"] = "b"
	fmt.Println(getMessage("%saccount_enquiry_error|param=%v", []interface{}{param}))
}
