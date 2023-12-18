package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingActivityLogsRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScalingActivityLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingActivityLogsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingActivityLogsRequest", string(data)}, " ")
}
