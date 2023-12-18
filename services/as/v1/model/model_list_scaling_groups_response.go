package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingGroupsResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingGroups  *[]ScalingGroups `json:"scaling_groups,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListScalingGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingGroupsResponse", string(data)}, " ")
}
