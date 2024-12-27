package main

type ResponseCode string

const (
	ResponseCodeSuccess ResponseCode = "0000"
	ResponseCodeFail    ResponseCode = "9999"
	ResponseCodeUnknown ResponseCode = "9998"
)

var AllResponseCode = []ResponseCode{
	ResponseCodeSuccess,
	ResponseCodeFail,
	ResponseCodeUnknown,
}

type Person struct {
	Name string
	Age  int
}

//func main() {
//	var people []*Person
//	fmt.Println(people == nil)
//	fmt.Println(len(people) == 0)
//}
