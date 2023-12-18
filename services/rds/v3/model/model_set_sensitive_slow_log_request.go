package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetSensitiveSlowLogRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Status string `json:"status"`
}

func (o SetSensitiveSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSensitiveSlowLogRequest struct{}"
	}

	return strings.Join([]string{"SetSensitiveSlowLogRequest", string(data)}, " ")
}
