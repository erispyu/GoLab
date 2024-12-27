package main

import "net/url"

func main() {
	requestUrl, _ := url.JoinPath("https://example.com", "/api/v1/endpoint")
	println(requestUrl)
	requestUrl, _ = url.JoinPath("https://example.com/", "/api/v1/endpoint")
	println(requestUrl)
	requestUrl, _ = url.JoinPath("https://example.com", "api/v1/endpoint")
	println(requestUrl)
}
