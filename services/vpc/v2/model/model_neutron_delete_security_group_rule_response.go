package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteSecurityGroupRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRuleResponse", string(data)}, " ")
}
