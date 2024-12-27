package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	RspBodyNIl = errors.New("resp body is nil")
)

type ApiClient struct {
	Ctx context.Context
}

func NewApiClient(ctx context.Context) *ApiClient {
	return &ApiClient{Ctx: ctx}
}

func (v ApiClient) GetToken() (*GetTokenResponse, error) {
	// req := &GetTokenRequest{
	// 	GrantType: "client_credentials",
	// }
	req := "grant_type=client_credentials"
	basicAuth := fmt.Sprintf("%s:%s", "mofA5VUqGrQRZbsihr2PsEriap0NYPYa", "DNja6hdqFYfRY03r")
	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(basicAuth)),
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	resp := new(GetTokenResponse)
	err := v.CallApi("/private/oauth2/v1/token", http.MethodPost, req, headers, resp)
	return resp, err
}

func (v ApiClient) ConvertCardNoToCardId(
	cardNumber string,
) (*ConvertCardNoToCardIdResponse, error) {
	req := &ConvertCardNoToCardIdRequest{
		CardNumber: cardNumber,
		RequestID:  uuid.NewV4().String(),
	}
	reqData, _ := json.Marshal(req)
	hashDigest := hmac.New(sha256.New, reqData).Sum(nil)
	headers := map[string]string{
		"user": "viettelpost",
		"hmac": string(hashDigest),
	}
	resp := new(ConvertCardNoToCardIdResponse)
	err := v.CallApi("/mbcardgw/internet/cardinfo/v1_0/generatetoken", http.MethodPost, req, headers, resp)
	return resp, err
}

func (v ApiClient) CallApi(
	endpoint string,
	method string,
	req interface{},
	headers map[string]string,
	resp interface{},
) error {
	var requestData []byte
	// if req type is string, convert it to []byte directly
	if reqStr, ok := req.(string); ok {
		requestData = []byte(reqStr)
	} else {
		var err error
		requestData, err = json.Marshal(req)
		if err != nil {
			return err
		}
	}

	httpRequest, err := v.buildHttpRequest(endpoint, method, requestData, headers)
	if err != nil {
		return err
	}

	client := v.getHttpClient()

	// send request

	httpResp, err := client.Do(httpRequest)

	if err != nil {
		return err
	}
	if resp == nil {
		return RspBodyNIl
	}

	// read response
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	defer func() {
		err := httpResp.Body.Close()
		if err != nil {
			return
		}
	}()

	return json.Unmarshal(respBody, resp)
}

func (v ApiClient) buildHttpRequest(
	endpoint string,
	method string,
	requestData []byte,
	headers map[string]string,
) (httpRequest *http.Request, err error) {
	baseUrl := "https://api-sandbox.mbbank.com.vn"
	fullUrl := path.Join(baseUrl, endpoint)

	httpRequest, err = http.NewRequest(method, fullUrl, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		httpRequest.Header.Set(k, v)
	}

	return httpRequest, err
}

func (v ApiClient) getHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 20 * time.Second,
	}
	return client
}
