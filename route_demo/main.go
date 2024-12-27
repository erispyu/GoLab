package main

import (
	"errors"
	"fmt"
	"sort"
)

type ProductChannel struct {
	ChannelId        int
	Priority         int
	Active           bool
	BiFastPayoutFail bool
}

type ProductChannelSliceDecrement []ProductChannel

func (p ProductChannelSliceDecrement) Len() int {
	return len(p)
}

func (p ProductChannelSliceDecrement) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var ChannelIdBiFastPayoutId = 8019400

func (p ProductChannelSliceDecrement) Less(i, j int) bool {
	return p[i].Priority > p[j].Priority
}

func getRouteByPriority(fallback, aePassed bool, productChannels []ProductChannel) (int, error) {
	if len(productChannels) == 0 {
		return 0, errors.New("no available route")
	}

	if len(productChannels) == 1 {
		if productChannels[0].ChannelId == ChannelIdBiFastPayoutId {
			if !aePassed {
				return 0, errors.New("no available route")
			}
		}
		return 0, nil
	}

	sort.Sort(ProductChannelSliceDecrement(productChannels))
	if productChannels[0].Priority == productChannels[1].Priority {
		return 0, errors.New("no highest priority channel")
	}
	for i, pc := range productChannels {
		if pc.Active {
			if pc.ChannelId == ChannelIdBiFastPayoutId && !aePassed {
				productChannels[i].BiFastPayoutFail = true
				if !fallback {
					return 0, errors.New("BI-Fast Error")
				}
				continue
			}
			return pc.ChannelId, nil
		}
		if i+1 < len(productChannels) && pc.Priority == productChannels[i+1].Priority {
			break
		}
		if !fallback {
			return pc.ChannelId, nil
		}
	}
	if productChannels[0].BiFastPayoutFail {
		return 0, errors.New("BI-Fast Error")
	}
	return productChannels[0].ChannelId, nil
}

func printRoute(route int, err error) {
	if route != 0 {
		fmt.Println("route:", route)
	} else {
		fmt.Println("err:", err)
	}
}

func main() {
	BIFastChannel := ProductChannel{
		ChannelId: ChannelIdBiFastPayoutId,
		Priority:  10,
		Active:    true,
	}
	OtherChannel := ProductChannel{
		ChannelId: 12,
		Priority:  9,
		Active:    true,
	}

	// Table Headers
	fmt.Println("Scenario #\tChannel Count\tFallback\tBI-Fast Status\tAE Status\tOther Channel Status\troute result")

	// Scenario #1: only 1 channel, fallback true, BI-Fast Active, AE Passed, Other Channel Active
	fmt.Println("Scenario #1:\tmore than 1 channel\tfallback true\tBI-Fast Active\tAE Passed\tOther Channel Active")
	printRoute(getRouteByPriority(true, true, []ProductChannel{BIFastChannel, OtherChannel}))

	// Scenario #1: more than 1 channel, fallback true, BI-Fast Active, AE Passed, Other Channel Active
	fmt.Println("Scenario #1:\tmore than 1 channel\tfallback true\tBI-Fast Active\tAE Passed\tOther Channel Active")
	printRoute(getRouteByPriority(true, true, []ProductChannel{BIFastChannel, OtherChannel}))

	// Scenario #2: more than 1 channel, fallback true, BI-Fast Active, AE Passed, Other Channel Inactive
	fmt.Println("Scenario #2:\tmore than 1 channel\tfallback true\tBI-Fast Active\tAE Passed\tOther Channel Inactive")
	OtherChannel.Active = false
	printRoute(getRouteByPriority(true, true, []ProductChannel{BIFastChannel, OtherChannel}))
}
