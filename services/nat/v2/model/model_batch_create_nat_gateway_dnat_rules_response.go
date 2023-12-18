package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateNatGatewayDnatRulesResponse struct {
	DnatRules      *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchCreateNatGatewayDnatRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNatGatewayDnatRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesResponse", string(data)}, " ")
}
