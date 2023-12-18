package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddServerNicOption struct {
	SubnetId string `json:"subnet_id"`

	SecurityGroups *[]ServerNicSecurityGroup `json:"security_groups,omitempty"`

	IpAddress *string `json:"ip_address,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6Bandwidth *Ipv6Bandwidth `json:"ipv6_bandwidth,omitempty"`
}

func (o BatchAddServerNicOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerNicOption struct{}"
	}

	return strings.Join([]string{"BatchAddServerNicOption", string(data)}, " ")
}
