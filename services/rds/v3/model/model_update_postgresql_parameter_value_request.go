package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdatePostgresqlParameterValueRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	Body *ModifyParamRequest `json:"body,omitempty"`
}

func (o UpdatePostgresqlParameterValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlParameterValueRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlParameterValueRequest", string(data)}, " ")
}
