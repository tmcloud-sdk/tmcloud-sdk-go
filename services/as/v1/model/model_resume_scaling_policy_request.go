package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResumeScalingPolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *ResumeScalingPolicyOption `json:"body,omitempty"`
}

func (o ResumeScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ResumeScalingPolicyRequest", string(data)}, " ")
}
