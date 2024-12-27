package main

import (
	"fmt"
	"reflect"
	"time"
)

// func main() {
// 	var passList, failList, stuckList []int
// 	returnNew(&passList, &failList, &stuckList)
// 	fmt.Println("passList:", passList)
// 	fmt.Println("failList:", failList)
// 	fmt.Println("stuckList:", stuckList)
// }
//
// func processSingle(job int) (pass, fail, stuck bool) {
// 	if job%3 == 0 {
// 		return true, false, false
// 	} else if job%3 == 1 {
// 		return false, true, false
// 	} else {
// 		return false, false, true
// 	}
// }
//
// func returnNew(passList, failList, stuckList *[]int) {
// 	for i := 0; i < 100; i++ {
// 		pass, fail, _ := processSingle(i)
// 		if pass {
// 			*passList = append(*passList, i)
// 		} else if fail {
// 			*failList = append(*failList, i)
// 		} else {
// 			*stuckList = append(*stuckList, i)
// 		}
// 	}
// 	fmt.Println("inner passList:", passList)
// 	fmt.Println("inner failList:", failList)
// 	fmt.Println("inner stuckList:", stuckList)
// }

func returnNew(list []int64) []int64 {
	return append(list, time.Now().Unix())
}

func returnOriginal(list []int64) {
	list = append(list, time.Now().Unix())
}

// func main() {
// 	var list []int64
// 	list = returnNew(list)
// 	fmt.Println(list)
// 	returnOriginal(list)
// 	fmt.Println(list)
// 	fmt.Println(strings.TrimSuffix("abc_123.txt.csv", ".csv"))
// }

func helper(s1, s2 []interface{}) []interface{} {
	var result []interface{}
	result = append(result, s1...)
	result = append(result, s2...)
	return result
}

type Channel struct {
	ChannelId uint32
	Name      string
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

func IsEmptyStruct(s interface{}) bool {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		if !IsZeroOfUnderlyingType(v.Field(i).Interface()) {
			return false
		}
	}
	return true
}

func main() {
	var result []*Channel
	result = append(result, &Channel{1, "Payoneer"})
	result = append(result, &Channel{2, "PingPong"})
	result = append(result, &Channel{})

	finalResult := make([]*Channel, 0, len(result))
	for _, res := range result {
		if res != nil && !IsEmptyStruct(res) {
			finalResult = append(finalResult, res)
		}
	}

	fmt.Println(finalResult)
}
