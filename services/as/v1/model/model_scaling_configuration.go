package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ScalingConfiguration struct {
	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`

	Tenant *string `json:"tenant,omitempty"`

	ScalingConfigurationName *string `json:"scaling_configuration_name,omitempty"`

	InstanceConfig *InstanceConfigResult `json:"instance_config,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	ScalingGroupId *string `json:"scaling_group_id,omitempty"`
}

func (o ScalingConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingConfiguration struct{}"
	}

	return strings.Join([]string{"ScalingConfiguration", string(data)}, " ")
}
