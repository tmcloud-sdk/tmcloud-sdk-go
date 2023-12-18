package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryUserDetailResp struct {
	Id *string `json:"id,omitempty"`

	Account *string `json:"account,omitempty"`

	Comment *string `json:"comment,omitempty"`

	IsTransfer *bool `json:"is_transfer,omitempty"`

	Privileges *string `json:"privileges,omitempty"`

	Password *string `json:"password,omitempty"`

	Roles *[]string `json:"roles,omitempty"`

	Selected *bool `json:"selected,omitempty"`

	NoPrivileges *string `json:"no_privileges,omitempty"`

	ParentAccount *string `json:"parent_account,omitempty"`

	NoParentAccount *string `json:"no_parent_account,omitempty"`
}

func (o QueryUserDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryUserDetailResp struct{}"
	}

	return strings.Join([]string{"QueryUserDetailResp", string(data)}, " ")
}
