package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UserAccountVo struct {
	Id string `json:"id"`

	Account string `json:"account"`

	Comment *string `json:"comment,omitempty"`

	IsTransfer bool `json:"is_transfer"`

	Privileges *[]string `json:"privileges,omitempty"`

	Password *string `json:"password,omitempty"`

	IsSetPassword *bool `json:"is_set_password,omitempty"`

	Roles []string `json:"roles"`

	Selected bool `json:"selected"`
}

func (o UserAccountVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAccountVo struct{}"
	}

	return strings.Join([]string{"UserAccountVo", string(data)}, " ")
}
