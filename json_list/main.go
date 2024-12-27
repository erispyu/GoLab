package main

import (
	"encoding/json"
	"fmt"
)

type PoolSize struct {
	ChannelId uint64 `json:"channel_id,omitempty"`
	Value     int    `json:"value,omitempty"`
}

func main() {
	configStr := `[{"channel_id":10015300,"value":5}]`
	var config []PoolSize
	err := json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}
