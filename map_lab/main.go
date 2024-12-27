package main

import "fmt"

func main() {
	// m := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }
	// fmt.Println(m)
	// TestMapReference(m)
	// fmt.Println(m)

	hashMap := map[string]bool{
		"a": true,
	}
	v1, _ := hashMap["a"]
	fmt.Println(v1)
	v2, _ := hashMap["b"]
	fmt.Println(v2)

	tableNameFormat := "recipient_info_tab_%08d"
	for i := 0; i < 1000; i++ {
		tableName := fmt.Sprintf(tableNameFormat, uint32(i))
		fmt.Println(tableName)
	}

	fmt.Println("Length of map:", len(hashMap))
}
