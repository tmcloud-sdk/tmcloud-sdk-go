package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ExecuteScalingPolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *ExecuteScalingPolicyOption `json:"body,omitempty"`
}

func (o ExecuteScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyRequest", string(data)}, " ")
}
