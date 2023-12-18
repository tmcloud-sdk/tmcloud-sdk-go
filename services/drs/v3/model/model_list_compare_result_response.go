package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListCompareResultResponse struct {
	JobId *string `json:"job_id,omitempty"`

	ObjectLevelCompareResults *ObjectCompareResult `json:"object_level_compare_results,omitempty"`

	LineCompareResults *LineCompareResult `json:"line_compare_results,omitempty"`

	ContentCompareResults *ContentCompareResult `json:"content_compare_results,omitempty"`

	CompareTaskListResults *CompareTaskListResult `json:"compare_task_list_results,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCompareResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompareResultResponse struct{}"
	}

	return strings.Join([]string{"ListCompareResultResponse", string(data)}, " ")
}
