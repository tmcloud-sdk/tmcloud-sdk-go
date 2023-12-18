package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationResponseApplyResults struct {
	InstanceId string `json:"instance_id"`

	InstanceName string `json:"instance_name"`

	RestartRequired bool `json:"restart_required"`

	Success bool `json:"success"`
}

func (o ApplyConfigurationResponseApplyResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationResponseApplyResults struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationResponseApplyResults", string(data)}, " ")
}
