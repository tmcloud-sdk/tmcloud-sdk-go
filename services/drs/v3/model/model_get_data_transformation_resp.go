package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GetDataTransformationResp struct {
	TotalCount *int64 `json:"total_count,omitempty"`

	FilterConditions *[]DataTransformationInfo `json:"filter_conditions,omitempty"`
}

func (o GetDataTransformationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDataTransformationResp struct{}"
	}

	return strings.Join([]string{"GetDataTransformationResp", string(data)}, " ")
}
