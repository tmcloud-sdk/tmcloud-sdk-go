package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowPostgresqlParamValueRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Name string `json:"name"`
}

func (o ShowPostgresqlParamValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPostgresqlParamValueRequest struct{}"
	}

	return strings.Join([]string{"ShowPostgresqlParamValueRequest", string(data)}, " ")
}
