package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewaySnatRuleOption struct {
	NatGatewayId string `json:"nat_gateway_id"`

	Cidr *string `json:"cidr,omitempty"`

	NetworkId *string `json:"network_id,omitempty"`

	Description *string `json:"description,omitempty"`

	SourceType *int32 `json:"source_type,omitempty"`

	FloatingIpId string `json:"floating_ip_id"`
}

func (o CreateNatGatewaySnatRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleOption", string(data)}, " ")
}
