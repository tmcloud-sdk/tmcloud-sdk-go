package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateL7PolicyRuleOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Type string `json:"type"`

	CompareType string `json:"compare_type"`

	Invert *bool `json:"invert,omitempty"`

	Key *string `json:"key,omitempty"`

	Value string `json:"value"`

	Conditions *[]CreateRuleCondition `json:"conditions,omitempty"`
}

func (o CreateL7PolicyRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyRuleOption struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyRuleOption", string(data)}, " ")
}
