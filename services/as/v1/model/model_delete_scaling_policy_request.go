package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteScalingPolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o DeleteScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingPolicyRequest", string(data)}, " ")
}
