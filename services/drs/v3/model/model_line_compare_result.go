package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LineCompareResult struct {
	CompareTaskId *string `json:"compare_task_id,omitempty"`

	LineCompareOverview *[]LineCompareResultOverview `json:"line_compare_overview,omitempty"`

	LineCompareOverviewCount *int32 `json:"line_compare_overview_count,omitempty"`

	LineCompareDetails *[]LineCompareResultDetails `json:"line_compare_details,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o LineCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResult struct{}"
	}

	return strings.Join([]string{"LineCompareResult", string(data)}, " ")
}
