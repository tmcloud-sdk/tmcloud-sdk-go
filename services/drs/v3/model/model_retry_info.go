package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RetryInfo struct {
	JobId string `json:"job_id"`

	IsSyncReEdit *bool `json:"is_sync_re_edit,omitempty"`
}

func (o RetryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryInfo struct{}"
	}

	return strings.Join([]string{"RetryInfo", string(data)}, " ")
}
