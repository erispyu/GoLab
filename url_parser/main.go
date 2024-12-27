package main

import (
	"net/url"
)

func getUrlParser(rawUrl string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	if parsedURL == nil || parsedURL.Host == "" || parsedURL.Path == "" {
		return parsedURL, err
	}
	return parsedURL, nil
}

func main() {
	rawUrl := "https://alphaapihub.cimbniaga.co.id:443/api-manager-external/alpha/v1.0/access-token/b2b"
	parsed, err := getUrlParser(rawUrl)
	if err != nil {
		panic(err)
	}
	if parsed != nil {
		println(parsed.Host)
		println(parsed.Path)
	}
}
