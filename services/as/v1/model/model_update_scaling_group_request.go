package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateScalingGroupRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *UpdateScalingGroupOption `json:"body,omitempty"`
}

func (o UpdateScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingGroupRequest", string(data)}, " ")
}
