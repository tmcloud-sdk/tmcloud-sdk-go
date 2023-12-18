package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationAsyncRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequest `json:"body,omitempty"`
}

func (o ApplyConfigurationAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationAsyncRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationAsyncRequest", string(data)}, " ")
}
