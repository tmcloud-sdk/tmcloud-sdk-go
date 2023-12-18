package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteL7RuleRequest struct {
	L7policyId string `json:"l7policy_id"`

	L7ruleId string `json:"l7rule_id"`
}

func (o DeleteL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7RuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7RuleRequest", string(data)}, " ")
}
