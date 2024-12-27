package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Status int64

const (
	StatusInitial Status = iota
	StatusPending
	StatusSuccess
	StatusFailed
)

var statusNames = map[Status]string{
	StatusInitial: "initial",
	StatusPending: "pending",
	StatusSuccess: "success",
	StatusFailed:  "failed",
}

func (p Status) Name() string {
	name, ok := statusNames[p]
	if !ok {
		return ""
	}
	return name
}

type Response struct {
	Status interface{}
}

func main() {
	// r1 := Response{Status: StatusInitial}
	// success := StatusSuccess
	// r2 := Response{Status: &success}
	//
	// rList := []Response{r1, r2}
	// for _, r := range rList {
	// 	status := r.Status
	// 	switch status := status.(type) {
	// 	case Status:
	// 		println(status.Name())
	// 	//case *Status:
	// 	//	println(status.Name())
	// 	default:
	// 		println("no status found")
	// 	}
	// }
	// testReceiver()
	// icOrCompanyIdRegExp := regexp.MustCompile(`(^[a-zA-Z]\d{9}$)|(^[a-zA-Z]{2}\d{8}$)|(^\d{8}$)`)
	// s := "S 125280813"
	// fmt.Println(icOrCompanyIdRegExp.MatchString(s))
	productConfigExtraData := spmjson.GetJsonReader().ParseJsonString(productConfigTab.ExtraData)
	fmt.Println(productConfigExtraData.Get("fallback_to_active_channel").ToBool())
}

func testReceiver() {
	var result string
	operator := func(receiver interface{}) (i interface{}, e error) {
		r := receiver.(*string)
		*r = "hello"
		return r, nil
	}
	err := ChangeReceiver(&result, operator)
	fmt.Printf("result: %s, err: %v\n", result, err)
}

type ChangeOperator func(receiver interface{}) (i interface{}, e error)

func ChangeReceiver(receiver interface{}, operator ChangeOperator) error {
	rv := reflect.ValueOf(receiver)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("receiver must be a non-nil pointer")
	}
	_, err := operator(receiver)
	return err
}
