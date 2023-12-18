package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyDbPwdResp struct {
	Id *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	EndPointType *string `json:"end_point_type,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ModifyDbPwdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbPwdResp struct{}"
	}

	return strings.Join([]string{"ModifyDbPwdResp", string(data)}, " ")
}
