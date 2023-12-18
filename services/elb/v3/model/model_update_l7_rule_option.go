package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7RuleOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	CompareType *string `json:"compare_type,omitempty"`

	Invert *bool `json:"invert,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`

	Conditions *[]UpdateRuleCondition `json:"conditions,omitempty"`
}

func (o UpdateL7RuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleOption struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleOption", string(data)}, " ")
}
