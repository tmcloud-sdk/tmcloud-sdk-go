package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateCompareTaskResponse struct {
	JobId *string `json:"job_id,omitempty"`

	ObjectLevelCompareCreateResult *CreateCompareTaskResult `json:"object_level_compare_create_result,omitempty"`

	DataLevelCompareCreateResult *CreateCompareTaskResult `json:"data_level_compare_create_result,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCompareTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskResponse", string(data)}, " ")
}
