package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingV2PoliciesResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingPolicies *[]ScalingPoliciesV2 `json:"scaling_policies,omitempty"`
	HttpStatusCode  int                  `json:"-"`
}

func (o ListScalingV2PoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingV2PoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListScalingV2PoliciesResponse", string(data)}, " ")
}
