package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AllowSqlserverDbUserPrivilegeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SqlserverGrantRequest `json:"body,omitempty"`
}

func (o AllowSqlserverDbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowSqlserverDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"AllowSqlserverDbUserPrivilegeRequest", string(data)}, " ")
}
