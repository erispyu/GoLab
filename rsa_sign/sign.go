package main

//
//import (
//	"GoLab/util"
//	"fmt"
//	"os"
//	"strings"
//	"time"
//)
//
//func signParams(api string, requestData *RequestData) (string, error) {
//	pairs := sortedKeyValuePairs(requestData)
//	signedStr := strings.Join(pairs, "&")
//	fmt.Println("Original JSON:", requestData)
//	fmt.Println("Sorted JSON:", signedStr)
//	privateKeyLocalPath := "/Users/chong.chen/myProjects/GoLab/rsa_sign/shopee.pri"
//
//	privateKeyData, err := os.ReadFile(privateKeyLocalPath)
//	if err != nil {
//		return "", err
//	}
//
//	return util.RsaSignWithPrivateKeyPem(api+"?"+signedStr, privateKeyData)
//}
//
//func testSign() {
//	api := "/api/mybank/enterprise/pay/qpayent/V1"
//	d := time.Unix(1621411200, 0)
//	timestamp := d.Format("2006-01-02 15:04:05")
//	bizContent := BizContent{
//		TransCode: "QPAYENT",
//		TranDate:  d.Format("20060102"),
//		TranTime:  d.Format("150405000"),
//		FSeqNo:    "abcd",
//		QryfSeqno: "abcdf",
//	}
//	requestData := &RequestData{
//		AppId:      "appid",
//		MsgId:      "1234",
//		SignType:   "RSA2",
//		Timestamp:  timestamp,
//		BizContent: util.SafeToJson(bizContent),
//	}
//	sign, err := signParams(api, requestData)
//	println(sign)
//	if err != nil {
//		println(err.Error())
//	}
//}
