package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CompareTaskListResult struct {
	CompareTaskList *[]CompareTaskList `json:"compare_task_list,omitempty"`

	CompareTaskListCount *int32 `json:"compare_task_list_count,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`
}

func (o CompareTaskListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTaskListResult struct{}"
	}

	return strings.Join([]string{"CompareTaskListResult", string(data)}, " ")
}
