package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingPoliciesResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingPolicies *[]ScalingV1PolicyDetail `json:"scaling_policies,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListScalingPoliciesResponse", string(data)}, " ")
}
