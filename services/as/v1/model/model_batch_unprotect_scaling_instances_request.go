package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUnprotectScalingInstancesRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchUnprotectInstancesOption `json:"body,omitempty"`
}

func (o BatchUnprotectScalingInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnprotectScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchUnprotectScalingInstancesRequest", string(data)}, " ")
}
