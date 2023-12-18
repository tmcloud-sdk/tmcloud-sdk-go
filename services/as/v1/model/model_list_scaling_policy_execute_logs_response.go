package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingPolicyExecuteLogsResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingPolicyExecuteLog *[]ScalingPolicyExecuteLogList `json:"scaling_policy_execute_log,omitempty"`
	HttpStatusCode          int                            `json:"-"`
}

func (o ListScalingPolicyExecuteLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPolicyExecuteLogsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingPolicyExecuteLogsResponse", string(data)}, " ")
}
