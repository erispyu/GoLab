package main

import (
	"fmt"
	"net/http"
)

type AccessTokenReq struct {
	GrantType string `json:"grantType,omitempty"`
}

type AccessTokenResp struct {
	ResponseCode    string `json:"responseCode,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
	AccessToken     string `json:"accessToken,omitempty"`
	TokenType       string `json:"tokenType,omitempty"`
	ExpiresIn       string `json:"expiresIn,omitempty"`
}

func (resp *AccessTokenResp) GetToken() string {
	if resp == nil {
		return ""
	}
	// For Open API MWS, the token would not contain the type prefix
	// Refer to: https://docs.google.com/document/d/10fqKv4ZkGj-t51c4sexxhVpeadmUAKnWdeCN6wAUUSw/edit 3.2.4.2.2
	return resp.AccessToken
}

type CreateOrderReq struct {
	PartnerReferenceNo string         `json:"partnerReferenceNo,omitempty"` // M
	CustomerNumber     string         `json:"customerNumber,omitempty"`     // M
	Amount             Amount         `json:"amount,omitempty"`             // M
	Notes              string         `json:"notes,omitempty"`              // O
	AdditionalInfo     AdditionalInfo `json:"additionalInfo,omitempty"`     // M
}

type CreateOrderResp struct {
	ResponseCode       string         `json:"responseCode,omitempty"`
	ResponseMessage    string         `json:"responseMessage,omitempty"`
	PartnerReferenceNo string         `json:"partnerReferenceNo,omitempty"`
	ReferenceNo        string         `json:"referenceNo,omitempty"`
	Amount             Amount         `json:"amount,omitempty"`
	AdditionalInfo     AdditionalInfo `json:"additionalInfo,omitempty"`
}

type QueryOrderReq struct {
	OriginalPartnerReferenceNo string         `json:"originalPartnerReferenceNo,omitempty"` // M
	ServiceCode                string         `json:"serviceCode,omitempty"`                // M
	AdditionalInfo             AdditionalInfo `json:"additionalInfo,omitempty"`             // M
}

type QueryOrderResp struct {
	ResponseCode               string         `json:"responseCode,omitempty"`
	ResponseMessage            string         `json:"responseMessage,omitempty"`
	OriginalPartnerReferenceNo string         `json:"originalPartnerReferenceNo,omitempty"`
	Amount                     Amount         `json:"amount,omitempty"`
	OriginalReferenceNo        string         `json:"originalReferenceNo,omitempty"`
	ServiceCode                string         `json:"serviceCode,omitempty"`
	LatestTransactionStatus    string         `json:"latestTransactionStatus,omitempty"`
	AdditionalInfo             AdditionalInfo `json:"additionalInfo,omitempty"`
}

type PreCheckReq struct {
	CustomerNumber string         `json:"customerNumber,omitempty"` // M
	Amount         Amount         `json:"amount,omitempty"`         // M
	AdditionalInfo AdditionalInfo `json:"additionalInfo,omitempty"` // M
}

type PreCheckResp struct {
	CustomerName   string         `json:"customerName,omitempty"`
	CustomerNumber string         `json:"customerNumber,omitempty"`
	Amount         Amount         `json:"amount,omitempty"`
	MinAmount      MinAmount      `json:"minAmount,omitempty"`
	FeeAmount      FeeAmount      `json:"feeAmount,omitempty"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo,omitempty"`
}

type BalanceReq struct {
	MerchantExtId string `json:"merchantExtId,omitempty"` // M
	StoreExtId    string `json:"storeExtId,omitempty"`    // O
}

type BalanceResp struct {
	ResponseCode     string               `json:"responseCode,omitempty"`
	ResponseMessage  string               `json:"responseMessage,omitempty"`
	AvailableBalance AvailableBalanceInfo `json:"availableBalance,omitempty"`
}

type AdditionalInfo struct {
	MerchantExtId           string            `json:"merchantExtId,omitempty"`       // M
	StoreExtId              string            `json:"storeExtId,omitempty"`          // M
	MerchantDisplayName     string            `json:"merchantDisplayName,omitempty"` // O
	ShopeeUserId            uint64            `json:"shopeeUserId,omitempty"`        // M
	ShopeePayUserId         uint64            `json:"shopeePayUserId,omitempty"`     // M
	UseCase                 int32             `json:"useCase,omitempty"`             // O
	Purpose                 int32             `json:"purpose,omitempty"`             // O
	UserRemarks             string            `json:"userRemarks,omitempty"`         // O
	SenderInfo              SenderInfo        `json:"senderInfo,omitempty"`          // M
	RecipientInfo           RecipientInfo     `json:"recipientInfo,omitempty"`       // M
	Destination             int32             `json:"destination,omitempty"`         // M, BANK_ACCOUNT = 2
	DisbursementQuota       DisbursementQuota `json:"disbursementQuota,omitempty"`   // O
	UserIdHash              string            `json:"userIdHash,omitempty"`
	LatestTransactionStatus string            `json:"latestTransactionStatus,omitempty"`
	CreateTime              string            `json:"createTime,omitempty"`
	UpdateTime              string            `json:"updateTime,omitempty"`
	FeeAmount               FeeAmountInfo     `json:"feeAmount,omitempty"`
	SpmSourceType           uint64            `json:"spmSourceType,omitempty"`
}

type SenderInfo struct {
	SenderName     string `json:"senderName,omitempty"`
	KycIcNo        string `json:"kycIcNo,omitempty"`
	KycPhoneNumber string `json:"kycPhoneNumber,omitempty"`
}

type RecipientInfo struct {
	BankAccountNumber string `json:"bankAccountNumber,omitempty"`
	BankAccountName   string `json:"bankAccountName,omitempty"`
	BankNameId        int32  `json:"bankNameId,omitempty"`
	BankAccountId     int64  `json:"bankAccountId,omitempty"`
	CardType          int32  `json:"cardType,omitempty"` // CARD_TYPE_BI_FAST_OB = 9
}

type Amount struct {
	Value    string `json:"value,omitempty"`    // M
	Currency string `json:"currency,omitempty"` // M
}

type FeeAmountInfo struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type AvailableBalanceInfo struct {
	Value         string `json:"value,omitempty"`
	Currency      string `json:"currency,omitempty"`
	MerchantExtId string `json:"merchantExtId,omitempty"`
	StoreExtId    string `json:"storeExtId,omitempty"`
}

type MinAmount struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}
type FeeAmount struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type DisbursementQuota struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// func main() {
// 	// data := `{"responseCode":"4013800","responseMessage":"Unauthorized.signature not matched"}`
// 	// var req *CreateOrderResp
// 	// err := json.Unmarshal([]byte(data), &req)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(req)
// 	// timestamp := time.Now().Unix()
// 	// incr := 324
// 	// s := fmt.Sprintf("%010d%08d", timestamp, incr)
// 	// fmt.Println(s)
// 	s := "Nguyễn Thị Minh Khôi"
// 	us := strings.ToUpper(s)
// 	fmt.Println(us)
// }

func handler(w http.ResponseWriter, r *http.Request) {
	// 返回指定的字符串
	responseString := ""
	fmt.Fprintln(w, responseString)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("服务器正在运行于 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
