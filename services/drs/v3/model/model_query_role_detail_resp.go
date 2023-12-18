package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryRoleDetailResp struct {
	Role *string `json:"role,omitempty"`

	Comment *string `json:"comment,omitempty"`

	IsTransfer *bool `json:"is_transfer,omitempty"`

	Privileges *string `json:"privileges,omitempty"`

	InheritsRoles *[]string `json:"inherits_roles,omitempty"`

	Selected *bool `json:"selected,omitempty"`
}

func (o QueryRoleDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRoleDetailResp struct{}"
	}

	return strings.Join([]string{"QueryRoleDetailResp", string(data)}, " ")
}
