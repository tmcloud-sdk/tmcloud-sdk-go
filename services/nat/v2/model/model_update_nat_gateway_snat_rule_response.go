package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewaySnatRuleResponse struct {
	SnatRule       *NatGatewayUpdateSnatRuleResponseBody `json:"snat_rule,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o UpdateNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleResponse", string(data)}, " ")
}
