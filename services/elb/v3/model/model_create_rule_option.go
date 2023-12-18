package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRuleOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	CompareType string `json:"compare_type"`

	Key *string `json:"key,omitempty"`

	Value string `json:"value"`

	ProjectId *string `json:"project_id,omitempty"`

	Type string `json:"type"`

	Invert *bool `json:"invert,omitempty"`

	Conditions *[]CreateRuleCondition `json:"conditions,omitempty"`
}

func (o CreateRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleOption struct{}"
	}

	return strings.Join([]string{"CreateRuleOption", string(data)}, " ")
}
