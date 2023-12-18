package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaServerInterfaceFixedIp struct {
	IpAddress string `json:"ip_address"`

	SubnetId string `json:"subnet_id"`
}

func (o NovaServerInterfaceFixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerInterfaceFixedIp struct{}"
	}

	return strings.Join([]string{"NovaServerInterfaceFixedIp", string(data)}, " ")
}
