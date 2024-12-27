package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func compressManageConfigValue(configValue string) string {
	compressed := strings.TrimSpace(configValue)
	var js json.RawMessage
	if json.Unmarshal([]byte(compressed), &js) != nil {
		// not json, return original string
		return compressed
	}
	// compress json
	compressedBytes, _ := json.Marshal(js)
	return string(compressedBytes)
}

func main() {
	str := `{
  "mandiri_traffic_per": 1,
  "seabank_traffic_per": 100,
  "traffic_per": 100,
  "name": 123
}`
	fmt.Println(compressManageConfigValue(str))
}
