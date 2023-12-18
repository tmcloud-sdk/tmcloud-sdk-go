package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRuleCondition struct {
	Key *string `json:"key,omitempty"`

	Value string `json:"value"`
}

func (o CreateRuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleCondition struct{}"
	}

	return strings.Join([]string{"CreateRuleCondition", string(data)}, " ")
}
