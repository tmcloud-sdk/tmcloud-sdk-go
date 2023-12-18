package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UserRoleVo struct {
	Role string `json:"role"`

	Comment *string `json:"comment,omitempty"`

	IsTransfer bool `json:"is_transfer"`

	Privileges string `json:"privileges"`

	InheritsRoles *[]string `json:"inherits_roles,omitempty"`

	Selected *bool `json:"selected,omitempty"`
}

func (o UserRoleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserRoleVo struct{}"
	}

	return strings.Join([]string{"UserRoleVo", string(data)}, " ")
}
