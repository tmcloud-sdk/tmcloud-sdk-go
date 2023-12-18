package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PauseScalingPolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *PauseScalingPolicyOption `json:"body,omitempty"`
}

func (o PauseScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"PauseScalingPolicyRequest", string(data)}, " ")
}
