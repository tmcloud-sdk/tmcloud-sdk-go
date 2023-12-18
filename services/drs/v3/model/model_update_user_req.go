package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateUserReq struct {
	JobId string `json:"job_id"`

	Password *string `json:"password,omitempty"`

	List *[]UserAccountVo `json:"list,omitempty"`

	UserRoles *[]UserRoleVo `json:"user_roles,omitempty"`

	IsSetPassword bool `json:"is_set_password"`

	IsMigrateUser bool `json:"is_migrate_user"`

	IsSyncObjectPrivilege *bool `json:"is_sync_object_privilege,omitempty"`
}

func (o UpdateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReq struct{}"
	}

	return strings.Join([]string{"UpdateUserReq", string(data)}, " ")
}
