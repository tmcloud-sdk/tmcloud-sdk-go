package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdatePostgresqlInstanceAliasRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *UpdateRdsInstanceAliasRequest `json:"body,omitempty"`
}

func (o UpdatePostgresqlInstanceAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlInstanceAliasRequest", string(data)}, " ")
}
