package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddScalingInstancesRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchAddInstancesOption `json:"body,omitempty"`
}

func (o BatchAddScalingInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchAddScalingInstancesRequest", string(data)}, " ")
}
