package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryRedisConfig struct {
	ParamValue *string `json:"param_value,omitempty"`

	ValueType *string `json:"value_type,omitempty"`

	ValueRange *string `json:"value_range,omitempty"`

	Description *string `json:"description,omitempty"`

	DefaultValue *string `json:"default_value,omitempty"`

	ParamName *string `json:"param_name,omitempty"`

	ParamId *string `json:"param_id,omitempty"`
}

func (o QueryRedisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRedisConfig struct{}"
	}

	return strings.Join([]string{"QueryRedisConfig", string(data)}, " ")
}
