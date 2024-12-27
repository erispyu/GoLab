package main

func main() {
	var data interface{}
	data = ""
	switch data.(type) {
	case string:
		println("string")
	default:
		println("default")
	}
}
