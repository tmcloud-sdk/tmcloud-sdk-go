package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRuleRequestBody struct {
	SecurityGroupRule *CreateSecurityGroupRuleOption `json:"security_group_rule"`
}

func (o CreateSecurityGroupRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleRequestBody", string(data)}, " ")
}
