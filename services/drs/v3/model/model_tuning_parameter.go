package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TuningParameter struct {
	ParamName *string `json:"param_name,omitempty"`

	ParamValue *string `json:"param_value,omitempty"`

	Availability *string `json:"availability,omitempty"`
}

func (o TuningParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TuningParameter struct{}"
	}

	return strings.Join([]string{"TuningParameter", string(data)}, " ")
}
