package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PauseScalingGroupRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *PauseScalingGroupOption `json:"body,omitempty"`
}

func (o PauseScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupRequest", string(data)}, " ")
}
