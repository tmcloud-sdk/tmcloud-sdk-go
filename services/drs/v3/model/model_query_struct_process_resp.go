package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryStructProcessResp struct {
	JobId string `json:"job_id"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`

	StructProcess *StructProcessResp `json:"struct_process,omitempty"`
}

func (o QueryStructProcessResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryStructProcessResp struct{}"
	}

	return strings.Join([]string{"QueryStructProcessResp", string(data)}, " ")
}
