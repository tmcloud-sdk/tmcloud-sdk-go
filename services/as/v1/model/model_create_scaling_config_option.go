package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingConfigOption struct {
	ScalingConfigurationName string `json:"scaling_configuration_name"`

	InstanceConfig *InstanceConfig `json:"instance_config"`
}

func (o CreateScalingConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigOption struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigOption", string(data)}, " ")
}
