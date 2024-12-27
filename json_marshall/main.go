package main

import (
	"encoding/json"
	"fmt"
)

type AccessTokenResp struct {
	ResponseCode    string `json:"responseCode,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
	AccessToken     string `json:"accessToken,omitempty"`
	TokenType       string `json:"tokenType,omitempty"`
	ExpiresIn       string `json:"expiresIn,omitempty"`
}

type GetTokenResponse struct {
	ResponseBase
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   string `json:"expiresIn"`
}

type ResponseCode = string

type ErrorDict struct {
	Error     string `json:"error,omitempty"` // should keep exactly the same with python enum name
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

type ResponseBase struct {
	ErrorDict
	ResponseCode    ResponseCode `json:"responseCode,omitempty"`
	ResponseMessage string       `json:"responseMessage,omitempty"`
}

func main() {
	resp := `{"responseCode":"2007300","responseMessage":"success","accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDbGllbnRJRCI6InNob3BlZS1tYXJrZXRwbGFjZSIsImV4cCI6MTcyMjQ5NjU0MzAwNiwiaWF0IjoxNzIyNDk1NjQzMDA2LCJpc3MiOiJzaG9wZWUtbWFya2V0cGxhY2UiLCJuYmYiOjE3MjI0OTU2NDMwMDZ9.QaJqCPbySNmS_8fBQAnpjDoiX2q_nzt-yuqnAbvbaKc","tokenType":"Bearer","expiresIn":"900"}`
	var r GetTokenResponse
	err := json.Unmarshal([]byte(resp), &r)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.AccessToken)

	rData, _ := json.Marshal(r)
	var r1 AccessTokenResp
	err = json.Unmarshal(rData, &r1)
	if err != nil {
		panic(err)
	}
	fmt.Println(r1.AccessToken)
}
