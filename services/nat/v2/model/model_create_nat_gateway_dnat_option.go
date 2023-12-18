package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewayDnatOption struct {
	Description *string `json:"description,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	NatGatewayId string `json:"nat_gateway_id"`

	InternalServicePort int32 `json:"internal_service_port"`

	FloatingIpId string `json:"floating_ip_id"`

	ExternalServicePort int32 `json:"external_service_port"`

	Protocol string `json:"protocol"`

	InternalServicePortRange *string `json:"internal_service_port_range,omitempty"`

	ExternalServicePortRange *string `json:"external_service_port_range,omitempty"`
}

func (o CreateNatGatewayDnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayDnatOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayDnatOption", string(data)}, " ")
}
