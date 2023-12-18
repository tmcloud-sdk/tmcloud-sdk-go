package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EnableConfigurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequest `json:"body,omitempty"`
}

func (o EnableConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableConfigurationRequest struct{}"
	}

	return strings.Join([]string{"EnableConfigurationRequest", string(data)}, " ")
}
