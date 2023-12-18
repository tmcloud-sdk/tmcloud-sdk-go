package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateL7RuleRequest struct {
	L7policyId string `json:"l7policy_id"`

	Body *CreateL7RuleRequestBody `json:"body,omitempty"`
}

func (o CreateL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7RuleRequest struct{}"
	}

	return strings.Join([]string{"CreateL7RuleRequest", string(data)}, " ")
}
