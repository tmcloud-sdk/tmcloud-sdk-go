package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AttachCallbackInstanceLifeCycleHookRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *CallbackLifeCycleHookOption `json:"body,omitempty"`
}

func (o AttachCallbackInstanceLifeCycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachCallbackInstanceLifeCycleHookRequest struct{}"
	}

	return strings.Join([]string{"AttachCallbackInstanceLifeCycleHookRequest", string(data)}, " ")
}
