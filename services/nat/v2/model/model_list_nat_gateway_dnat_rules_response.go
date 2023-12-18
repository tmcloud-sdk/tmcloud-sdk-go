package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListNatGatewayDnatRulesResponse struct {
	DnatRules      *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListNatGatewayDnatRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayDnatRulesResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewayDnatRulesResponse", string(data)}, " ")
}
