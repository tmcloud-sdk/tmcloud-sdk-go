package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteNatGatewayDnatRuleRequest struct {
	NatGatewayId string `json:"nat_gateway_id"`

	DnatRuleId string `json:"dnat_rule_id"`
}

func (o DeleteNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayDnatRuleRequest", string(data)}, " ")
}
