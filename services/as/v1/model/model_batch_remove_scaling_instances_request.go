package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchRemoveScalingInstancesRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchRemoveInstancesOption `json:"body,omitempty"`
}

func (o BatchRemoveScalingInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveScalingInstancesRequest", string(data)}, " ")
}
