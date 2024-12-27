package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func HmacSha256Base64(plainText string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(plainText))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func main() {
	clientKey := "Fkitc11XkzcziKPYoeTbwj1Eoc3xTiCsfIdWxUC8cFhoA2NxP2GNiHdWu81iywxp"
	requestStr := "{\n    \"request_id\":\"1697080050972751018-account-gateway\",\n    \"user_id\": 258167630\n}"

	var req interface{}
	err := json.Unmarshal([]byte(requestStr), &req)
	if err != nil {
		return
	}

	requestData, _ := json.Marshal(req)
	requestStr = string(requestData)
	fmt.Println(requestStr)

	sign := HmacSha256Base64(requestStr, clientKey)
	fmt.Println(sign)
}
