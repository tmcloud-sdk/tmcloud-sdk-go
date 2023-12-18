package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EnableConfigurationResponse struct {
	ConfigurationId *string `json:"configuration_id,omitempty"`

	ConfigurationName *string `json:"configuration_name,omitempty"`

	Success *bool `json:"success,omitempty"`

	ApplyResults   *[]ApplyConfigurationResponseApplyResults `json:"apply_results,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o EnableConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableConfigurationResponse struct{}"
	}

	return strings.Join([]string{"EnableConfigurationResponse", string(data)}, " ")
}
