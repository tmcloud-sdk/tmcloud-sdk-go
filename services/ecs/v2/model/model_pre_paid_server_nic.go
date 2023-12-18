package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PrePaidServerNic struct {
	SubnetId string `json:"subnet_id"`

	IpAddress *string `json:"ip_address,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6Bandwidth *PrePaidServerIpv6Bandwidth `json:"ipv6_bandwidth,omitempty"`
}

func (o PrePaidServerNic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerNic struct{}"
	}

	return strings.Join([]string{"PrePaidServerNic", string(data)}, " ")
}
