package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResumeScalingGroupRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *ResumeScalingGroupOption `json:"body,omitempty"`
}

func (o ResumeScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"ResumeScalingGroupRequest", string(data)}, " ")
}
