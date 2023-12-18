package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingGroupRequest struct {
	Body *CreateScalingGroupOption `json:"body,omitempty"`
}

func (o CreateScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupRequest", string(data)}, " ")
}
