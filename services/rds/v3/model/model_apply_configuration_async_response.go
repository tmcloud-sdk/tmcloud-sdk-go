package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationAsyncResponse struct {
	ConfigurationId *string `json:"configuration_id,omitempty"`

	ConfigurationName *string `json:"configuration_name,omitempty"`

	Success *bool `json:"success,omitempty"`

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyConfigurationAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationAsyncResponse struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationAsyncResponse", string(data)}, " ")
}
