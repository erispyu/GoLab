package main

import "fmt"

type ProductChannelConfigTab struct {
	ProductChannelConfigId uint64 `json:"product_channel_config_id"`
	ProductConfigId        uint64 `json:"product_config_id"`
	ChannelId              uint64 `json:"channel_id"`
	ChannelAccountId       uint64 `json:"channel_account_id"`
	Priority               uint32 `json:"priority"`
}

// func (t *ProductChannelConfigTab) String() string {
// 	return fmt.Sprintf(
// 		"{ProductChannelConfigId: %d, "+
// 			"ProductConfigId: %d, "+
// 			"ChannelId: %d, "+
// 			"ChannelAccountId: %d, "+
// 			"Priority: %d, ",
// 		t.ProductChannelConfigId,
// 		t.ProductConfigId,
// 		t.ChannelId,
// 		t.ChannelAccountId,
// 		t.Priority,
// 	)
// }

func main() {
	c1 := ProductChannelConfigTab{
		ProductChannelConfigId: 1,
		ProductConfigId:        1,
		ChannelId:              1,
		ChannelAccountId:       1,
		Priority:               1,
	}
	c2 := ProductChannelConfigTab{
		ProductChannelConfigId: 2,
		ProductConfigId:        2,
		ChannelId:              2,
		ChannelAccountId:       2,
		Priority:               2,
	}
	clist := []ProductChannelConfigTab{c1, c2}
	fmt.Printf("%+v\n", clist)
}
