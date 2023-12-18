package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingNotificationRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *CreateNotificationOption `json:"body,omitempty"`
}

func (o CreateScalingNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingNotificationRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingNotificationRequest", string(data)}, " ")
}
