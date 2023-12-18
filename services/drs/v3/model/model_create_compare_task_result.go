package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateCompareTaskResult struct {
	CompareTaskId *string `json:"compare_task_id,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o CreateCompareTaskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskResult struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskResult", string(data)}, " ")
}
