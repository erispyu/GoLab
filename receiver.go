package main

import (
	"fmt"
	"reflect"
)

type Operator func(receiver interface{}) (i interface{}, e error)

func Handle(receiver interface{}, f Operator) error {
	rv := reflect.ValueOf(receiver)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("receiver must be a non-nil pointer")
	}
	_, err := f(receiver)
	return err
}

func main() {
	type Data struct {
		value bool
	}
	var val bool
	var processor Operator = func(receiver interface{}) (i interface{}, e error) {
		val = true
		return &Data{value: true}, nil
	}
	result := &Data{}
	err := Handle(result, processor)
	fmt.Println(result.value)
	fmt.Println(val)
	fmt.Println(err)
}
