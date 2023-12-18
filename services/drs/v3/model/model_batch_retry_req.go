package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchRetryReq struct {
	Jobs []RetryInfo `json:"jobs"`
}

func (o BatchRetryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRetryReq struct{}"
	}

	return strings.Join([]string{"BatchRetryReq", string(data)}, " ")
}
