package main

import (
	"errors"
	"fmt"
)

type BiFastPayoutAeError struct{}

func (e BiFastPayoutAeError) Error() string {
	return "BiFastPayoutAeError"
}

func main() {
	err := test()
	var biFastPayoutAeError BiFastPayoutAeError
	fmt.Println(errors.As(err, &biFastPayoutAeError))
}

func test() error {
	return BiFastPayoutAeError{}
}
