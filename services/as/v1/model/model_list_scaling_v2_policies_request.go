package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingV2PoliciesRequest struct {
	ScalingResourceId string `json:"scaling_resource_id"`

	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyType *string `json:"scaling_policy_type,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScalingV2PoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingV2PoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListScalingV2PoliciesRequest", string(data)}, " ")
}
