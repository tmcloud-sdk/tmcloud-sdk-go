package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingInstancesResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingGroupInstances *[]ScalingGroupInstance `json:"scaling_group_instances,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ListScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListScalingInstancesResponse", string(data)}, " ")
}
