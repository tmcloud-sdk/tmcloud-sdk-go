package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyInstancePasswordBody struct {
	OldPassword *string `json:"old_password,omitempty"`

	NewPassword *string `json:"new_password,omitempty"`
}

func (o ModifyInstancePasswordBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstancePasswordBody struct{}"
	}

	return strings.Join([]string{"ModifyInstancePasswordBody", string(data)}, " ")
}
