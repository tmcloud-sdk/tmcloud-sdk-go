package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceParamHistoriesRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ParamName *string `json:"param_name,omitempty"`
}

func (o ListInstanceParamHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceParamHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceParamHistoriesRequest", string(data)}, " ")
}
