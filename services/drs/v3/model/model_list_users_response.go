package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListUsersResponse struct {
	JobId *string `json:"job_id,omitempty"`

	IsGlobalPassword *string `json:"is_global_password,omitempty"`

	Message *string `json:"message,omitempty"`

	UserList *[]QueryUserDetailResp `json:"user_list,omitempty"`

	RolesList *[]QueryRoleDetailResp `json:"roles_list,omitempty"`

	IsSuccess      *bool `json:"is_success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
