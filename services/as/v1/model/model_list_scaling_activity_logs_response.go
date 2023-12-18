package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingActivityLogsResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingActivityLog *[]ScalingActivityLogList `json:"scaling_activity_log,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o ListScalingActivityLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingActivityLogsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingActivityLogsResponse", string(data)}, " ")
}
