package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowScalingGroupResponse struct {
	ScalingGroup   *ScalingGroups `json:"scaling_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowScalingGroupResponse", string(data)}, " ")
}
