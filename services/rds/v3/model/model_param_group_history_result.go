package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ParamGroupHistoryResult struct {
	ParameterName *string `json:"parameter_name,omitempty"`

	OldValue *string `json:"old_value,omitempty"`

	NewValue *string `json:"new_value,omitempty"`

	UpdateResult *string `json:"update_result,omitempty"`

	Applied *bool `json:"applied,omitempty"`

	UpdateTime *string `json:"update_time,omitempty"`

	ApplyTime *string `json:"apply_time,omitempty"`
}

func (o ParamGroupHistoryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupHistoryResult struct{}"
	}

	return strings.Join([]string{"ParamGroupHistoryResult", string(data)}, " ")
}
