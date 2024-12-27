package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	configStr1 := `{"total_switch":true}`
	configStr2 := `{"8015700":{"user_ids":[1,2,3],"enabled":true}}`
	fmt.Printf("Result: %v\n", check(configStr1))
	fmt.Println(check(configStr2))
}

func check(configStr string) bool {
	controlConfig := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(configStr), &controlConfig)
	if err != nil {
		return false
	}
	if _, ok := controlConfig["total_switch"]; ok {
		switch totalSwitch := controlConfig["total_switch"].(type) {
		case bool:
			if totalSwitch {
				return true
			}
		default:
			fmt.Printf("total_switch invalid type: %T\n", controlConfig["total_switch"])
		}
	}
	return false
}
