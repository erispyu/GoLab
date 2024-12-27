package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/samber/lo"
)

func NumToStr(value interface{}) string {
	o := reflect.ValueOf(value)
	kind := o.Kind()
	switch kind {
	case reflect.String:
		return o.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(o.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(o.Uint(), 10)
	}
	return ""
}

type BankAccount struct {
	Id *uint32 `json:"id,omitempty"`
}

func main() {
	ba := BankAccount{Id: lo.ToPtr(uint32(1111))}
	fmt.Println(NumToStr(ba.Id))
	fmt.Println(NumToStr(*ba.Id))
}
