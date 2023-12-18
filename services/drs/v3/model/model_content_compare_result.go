package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ContentCompareResult struct {
	CompareTaskId string `json:"compare_task_id"`

	ContentCompareOverview *[]ContentCompareResultOverview `json:"content_compare_overview,omitempty"`

	ContentCompareOverviewCount *int32 `json:"content_compare_overview_count,omitempty"`

	ContentCompareDetails *[]ContentCompareResultDetails `json:"content_compare_details,omitempty"`

	ContentCompareDiffs *[]ContentCompareResultDiffs `json:"content_compare_diffs,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ContentCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResult struct{}"
	}

	return strings.Join([]string{"ContentCompareResult", string(data)}, " ")
}
