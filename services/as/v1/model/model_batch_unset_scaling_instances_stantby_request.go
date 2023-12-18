package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUnsetScalingInstancesStantbyRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchExitStandByInstancesOption `json:"body,omitempty"`
}

func (o BatchUnsetScalingInstancesStantbyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnsetScalingInstancesStantbyRequest struct{}"
	}

	return strings.Join([]string{"BatchUnsetScalingInstancesStantbyRequest", string(data)}, " ")
}
