package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaServerInterfaceDetail struct {
	FixedIps []NovaServerInterfaceFixedIp `json:"fixed_ips"`

	MacAddr string `json:"mac_addr"`

	NetId string `json:"net_id"`

	PortId string `json:"port_id"`

	PortState string `json:"port_state"`
}

func (o NovaServerInterfaceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerInterfaceDetail struct{}"
	}

	return strings.Join([]string{"NovaServerInterfaceDetail", string(data)}, " ")
}
