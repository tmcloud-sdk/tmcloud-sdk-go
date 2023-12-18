package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ScalingInstance struct {
	InstanceName *string `json:"instance_name,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	FailedDetails *string `json:"failed_details,omitempty"`

	InstanceConfig *string `json:"instance_config,omitempty"`
}

func (o ScalingInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingInstance struct{}"
	}

	return strings.Join([]string{"ScalingInstance", string(data)}, " ")
}
