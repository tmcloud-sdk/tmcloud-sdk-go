package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListNatGatewaySnatRulesResponse struct {
	SnatRules      *[]NatGatewaySnatRuleResponseBody `json:"snat_rules,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListNatGatewaySnatRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaySnatRulesResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewaySnatRulesResponse", string(data)}, " ")
}
