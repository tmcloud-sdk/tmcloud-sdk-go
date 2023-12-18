package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRuleOption struct {
	SecurityGroupId string `json:"security_group_id"`

	Description *string `json:"description,omitempty"`

	Direction string `json:"direction"`

	Ethertype *string `json:"ethertype,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	PortRangeMin *int32 `json:"port_range_min,omitempty"`

	PortRangeMax *int32 `json:"port_range_max,omitempty"`

	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`

	RemoteGroupId *string `json:"remote_group_id,omitempty"`
}

func (o CreateSecurityGroupRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleOption", string(data)}, " ")
}
