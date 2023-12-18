package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InterfaceAttachment struct {
	FixedIps *[]ServerInterfaceFixedIp `json:"fixed_ips,omitempty"`

	MacAddr *string `json:"mac_addr,omitempty"`

	NetId *string `json:"net_id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	PortState *string `json:"port_state,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`

	DriverMode *string `json:"driver_mode,omitempty"`

	MinRate *int32 `json:"min_rate,omitempty"`

	MultiqueueNum *int32 `json:"multiqueue_num,omitempty"`

	PciAddress *string `json:"pci_address,omitempty"`
}

func (o InterfaceAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfaceAttachment struct{}"
	}

	return strings.Join([]string{"InterfaceAttachment", string(data)}, " ")
}
