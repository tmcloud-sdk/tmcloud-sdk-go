package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EniNetworkUpdate struct {
	Subnets *[]NetworkSubnet `json:"subnets,omitempty"`
}

func (o EniNetworkUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EniNetworkUpdate struct{}"
	}

	return strings.Join([]string{"EniNetworkUpdate", string(data)}, " ")
}
