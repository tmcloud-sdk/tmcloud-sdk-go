package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetDatabaseUserPrivilegeReqV3 struct {
	AllUsers bool `json:"all_users"`

	UserName *string `json:"user_name,omitempty"`

	Readonly bool `json:"readonly"`
}

func (o SetDatabaseUserPrivilegeReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDatabaseUserPrivilegeReqV3 struct{}"
	}

	return strings.Join([]string{"SetDatabaseUserPrivilegeReqV3", string(data)}, " ")
}
