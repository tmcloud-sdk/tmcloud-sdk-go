package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteScalingConfigOption struct {
	ScalingConfigurationId []string `json:"scaling_configuration_id"`
}

func (o BatchDeleteScalingConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingConfigOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingConfigOption", string(data)}, " ")
}
