package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSetScalingInstancesStandbyRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchEnterStandbyInstancesOption `json:"body,omitempty"`
}

func (o BatchSetScalingInstancesStandbyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetScalingInstancesStandbyRequest struct{}"
	}

	return strings.Join([]string{"BatchSetScalingInstancesStandbyRequest", string(data)}, " ")
}
