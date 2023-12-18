package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResetServerPasswordOption struct {
	NewPassword string `json:"new_password"`

	IsCheckPassword *bool `json:"is_check_password,omitempty"`
}

func (o ResetServerPasswordOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordOption struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordOption", string(data)}, " ")
}
