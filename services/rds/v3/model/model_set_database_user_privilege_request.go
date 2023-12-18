package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetDatabaseUserPrivilegeRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetDatabaseUserPrivilegeReqV3 `json:"body,omitempty"`
}

func (o SetDatabaseUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDatabaseUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"SetDatabaseUserPrivilegeRequest", string(data)}, " ")
}
