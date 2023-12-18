package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaAttachInterfaceOption struct {
	FixedIps *[]NovaAttachInterfaceFixedIp `json:"fixed_ips,omitempty"`

	NetId *string `json:"net_id,omitempty"`

	PortId *string `json:"port_id,omitempty"`
}

func (o NovaAttachInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceOption struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceOption", string(data)}, " ")
}
