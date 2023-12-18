package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryRpoAndRtoResp struct {
	JobId *string `json:"job_id,omitempty"`

	RpoInfo *RpoAndRtoInfo `json:"rpo_info,omitempty"`

	RtoInfo *RpoAndRtoInfo `json:"rto_info,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o QueryRpoAndRtoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRpoAndRtoResp struct{}"
	}

	return strings.Join([]string{"QueryRpoAndRtoResp", string(data)}, " ")
}
