package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ImportSmnResp struct {
	Id *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ImportSmnResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportSmnResp struct{}"
	}

	return strings.Join([]string{"ImportSmnResp", string(data)}, " ")
}
