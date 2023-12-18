package region

import (
	"fmt"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/region"
)

var (
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://nat.my-kualalumpur-1.alphaedge.tmone.com.my")
)

var staticFields = map[string]*region.Region{
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
}

var provider = region.DefaultProviderChain("NAT")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
