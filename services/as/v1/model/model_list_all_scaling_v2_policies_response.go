package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAllScalingV2PoliciesResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingPolicies *[]ScalingAllPolicyDetail `json:"scaling_policies,omitempty"`
	HttpStatusCode  int                       `json:"-"`
}

func (o ListAllScalingV2PoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllScalingV2PoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAllScalingV2PoliciesResponse", string(data)}, " ")
}
