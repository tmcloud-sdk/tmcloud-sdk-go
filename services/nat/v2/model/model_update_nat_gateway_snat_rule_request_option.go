package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewaySnatRuleRequestOption struct {
	SnatRule *UpdateNatGatewaySnatRuleOption `json:"snat_rule"`
}

func (o UpdateNatGatewaySnatRuleRequestOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleRequestOption struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleRequestOption", string(data)}, " ")
}
