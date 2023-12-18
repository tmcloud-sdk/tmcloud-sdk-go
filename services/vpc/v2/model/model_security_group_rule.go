package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SecurityGroupRule struct {
	Id string `json:"id"`

	Description string `json:"description"`

	SecurityGroupId string `json:"security_group_id"`

	Direction string `json:"direction"`

	Ethertype string `json:"ethertype"`

	Protocol string `json:"protocol"`

	PortRangeMin int32 `json:"port_range_min"`

	PortRangeMax int32 `json:"port_range_max"`

	RemoteIpPrefix string `json:"remote_ip_prefix"`

	RemoteGroupId string `json:"remote_group_id"`

	TenantId string `json:"tenant_id"`
}

func (o SecurityGroupRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupRule struct{}"
	}

	return strings.Join([]string{"SecurityGroupRule", string(data)}, " ")
}
