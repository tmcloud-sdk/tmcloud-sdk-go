package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateScalingPolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *UpdateScalingPolicyOption `json:"body,omitempty"`
}

func (o UpdateScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingPolicyRequest", string(data)}, " ")
}
