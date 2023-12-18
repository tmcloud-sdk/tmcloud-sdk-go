package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ObjectCompareResult struct {
	CompareTaskId string `json:"compare_task_id"`

	ObjectCompareOverview *[]ObjectCompareResultOverview `json:"object_compare_overview,omitempty"`

	ObjectCompareDetails map[string][]ObjectCompareResultDetails `json:"object_compare_details,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ObjectCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectCompareResult struct{}"
	}

	return strings.Join([]string{"ObjectCompareResult", string(data)}, " ")
}
