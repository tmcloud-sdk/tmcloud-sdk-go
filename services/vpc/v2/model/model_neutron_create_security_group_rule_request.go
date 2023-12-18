package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateSecurityGroupRuleRequest struct {
	Body *NeutronCreateSecurityGroupRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleRequest", string(data)}, " ")
}
