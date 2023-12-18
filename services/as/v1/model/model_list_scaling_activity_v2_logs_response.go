package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingActivityV2LogsResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingActivityLog *[]ScalingActivityLogV2 `json:"scaling_activity_log,omitempty"`
	HttpStatusCode     int                     `json:"-"`
}

func (o ListScalingActivityV2LogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingActivityV2LogsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingActivityV2LogsResponse", string(data)}, " ")
}
