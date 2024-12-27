package main

import (
	"context"
	"fmt"
)

func testToken() {
	ctx := context.Background()
	apiClient := NewApiClient(ctx)
	tokenResp, err := apiClient.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tokenResp)
}

func testConvertCardNoToCardId() {
	ctx := context.Background()
	apiClient := NewApiClient(ctx)
	resp, err := apiClient.ConvertCardNoToCardId("9704060129837294")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
