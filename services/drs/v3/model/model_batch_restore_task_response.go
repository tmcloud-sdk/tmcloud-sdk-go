package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchRestoreTaskResponse struct {
	Results *[]RetryTaskResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchRestoreTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchRestoreTaskResponse", string(data)}, " ")
}
