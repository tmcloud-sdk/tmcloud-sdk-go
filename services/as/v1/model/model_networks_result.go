package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NetworksResult struct {
	Id *string `json:"id,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6Bandwidth *Ipv6Bandwidth `json:"ipv6_bandwidth,omitempty"`

	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`
}

func (o NetworksResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworksResult struct{}"
	}

	return strings.Join([]string{"NetworksResult", string(data)}, " ")
}
