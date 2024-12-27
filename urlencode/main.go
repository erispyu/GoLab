package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := url.Values{}
	u.Add("app_id", "123")
	u.Add("msg_id", "456")
	u.Add("sign_type", "789")
	u.Add("timestamp", "2023-01-01")
	u.Add("biz_content", "hello world")
	fmt.Println(u.Encode())
}
