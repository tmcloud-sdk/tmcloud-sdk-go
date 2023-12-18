package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PauseJobResp struct {
	Id string `json:"id"`

	Status string `json:"status"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PauseJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseJobResp struct{}"
	}

	return strings.Join([]string{"PauseJobResp", string(data)}, " ")
}
