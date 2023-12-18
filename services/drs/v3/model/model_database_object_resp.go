package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DatabaseObjectResp struct {
	JobId *string `json:"job_id,omitempty"`

	Status *bool `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DatabaseObjectResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectResp struct{}"
	}

	return strings.Join([]string{"DatabaseObjectResp", string(data)}, " ")
}
