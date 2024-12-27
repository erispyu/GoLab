package main

import "fmt"

type ApiClientInterface interface {
	CallApi()
}

type ApiClientImpl struct {
	clientName string
}

func (c *ApiClientImpl) CallApi() {
	fmt.Printf("%s is calling api\n", c.clientName)
}

func main() {
	alice := &ApiClientImpl{clientName: "Alice"}
	alice.CallApi()

	var bob ApiClientInterface
	bob = &ApiClientImpl{clientName: "Bob"}
	bob.CallApi()
}
