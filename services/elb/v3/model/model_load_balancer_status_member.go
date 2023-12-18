package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusMember struct {
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	Address *string `json:"address,omitempty"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	Id *string `json:"id,omitempty"`

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusMember struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusMember", string(data)}, " ")
}
