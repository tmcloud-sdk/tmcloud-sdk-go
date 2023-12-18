package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateLifyCycleHookRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *CreateLifeCycleHookOption `json:"body,omitempty"`
}

func (o CreateLifyCycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLifyCycleHookRequest struct{}"
	}

	return strings.Join([]string{"CreateLifyCycleHookRequest", string(data)}, " ")
}
