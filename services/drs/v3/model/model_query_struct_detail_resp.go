package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryStructDetailResp struct {
	JobId string `json:"job_id"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`

	StructDetail *QueryFlowCompareDataResp `json:"struct_detail,omitempty"`
}

func (o QueryStructDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryStructDetailResp struct{}"
	}

	return strings.Join([]string{"QueryStructDetailResp", string(data)}, " ")
}
