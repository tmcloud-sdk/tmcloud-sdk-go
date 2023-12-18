package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AllowDbPrivilegeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *PostgresqlGrantRequest `json:"body,omitempty"`
}

func (o AllowDbPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"AllowDbPrivilegeRequest", string(data)}, " ")
}
