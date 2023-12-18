package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateLifeCycleHookRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	LifecycleHookName string `json:"lifecycle_hook_name"`

	Body *UpdateLifeCycleHookOption `json:"body,omitempty"`
}

func (o UpdateLifeCycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLifeCycleHookRequest struct{}"
	}

	return strings.Join([]string{"UpdateLifeCycleHookRequest", string(data)}, " ")
}
