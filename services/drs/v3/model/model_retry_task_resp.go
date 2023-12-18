package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RetryTaskResp struct {
	Id string `json:"id"`

	Status string `json:"status"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o RetryTaskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskResp struct{}"
	}

	return strings.Join([]string{"RetryTaskResp", string(data)}, " ")
}
