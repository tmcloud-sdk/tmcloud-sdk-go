package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingGroupResponse struct {
	ScalingGroupId *string `json:"scaling_group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupResponse", string(data)}, " ")
}
