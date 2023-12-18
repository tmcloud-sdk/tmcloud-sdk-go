package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewaySnatRuleOption struct {
	NatGatewayId string `json:"nat_gateway_id"`

	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateNatGatewaySnatRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleOption struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleOption", string(data)}, " ")
}
