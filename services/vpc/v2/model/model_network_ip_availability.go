package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NetworkIpAvailability struct {
	NetworkId string `json:"network_id"`

	NetworkName string `json:"network_name"`

	TenantId string `json:"tenant_id"`

	TotalIps int32 `json:"total_ips"`

	UsedIps int32 `json:"used_ips"`

	SubnetIpAvailability []SubnetIpAvailability `json:"subnet_ip_availability"`
}

func (o NetworkIpAvailability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkIpAvailability struct{}"
	}

	return strings.Join([]string{"NetworkIpAvailability", string(data)}, " ")
}
