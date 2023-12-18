package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SubnetIpAvailability struct {
	UsedIps int32 `json:"used_ips"`

	SubnetId string `json:"subnet_id"`

	SubnetName string `json:"subnet_name"`

	IpVersion int32 `json:"ip_version"`

	Cidr string `json:"cidr"`

	TotalIps int32 `json:"total_ips"`
}

func (o SubnetIpAvailability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetIpAvailability struct{}"
	}

	return strings.Join([]string{"SubnetIpAvailability", string(data)}, " ")
}
