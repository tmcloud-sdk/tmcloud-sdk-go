package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchPauseScalingPoliciesRequest struct {
	Body *BatchPauseScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchPauseScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesRequest", string(data)}, " ")
}
