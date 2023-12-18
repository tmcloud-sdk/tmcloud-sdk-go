package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingConfigsResponse struct {
	TotalNumber *int32 `json:"total_number,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ScalingConfigurations *[]ScalingConfiguration `json:"scaling_configurations,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ListScalingConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingConfigsResponse", string(data)}, " ")
}
