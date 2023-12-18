package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ExtraDhcpOpt struct {
	OptName *string `json:"opt_name,omitempty"`

	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOpt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraDhcpOpt struct{}"
	}

	return strings.Join([]string{"ExtraDhcpOpt", string(data)}, " ")
}
