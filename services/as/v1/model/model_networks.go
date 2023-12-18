package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Networks struct {
	Id string `json:"id"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6Bandwidth *Ipv6Bandwidth `json:"ipv6_bandwidth,omitempty"`

	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`
}

func (o Networks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Networks struct{}"
	}

	return strings.Join([]string{"Networks", string(data)}, " ")
}
