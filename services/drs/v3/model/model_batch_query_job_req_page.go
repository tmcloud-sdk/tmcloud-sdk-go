package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchQueryJobReqPage struct {
	Jobs []string `json:"jobs"`

	PageReq *PageReq `json:"page_req,omitempty"`
}

func (o BatchQueryJobReqPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryJobReqPage struct{}"
	}

	return strings.Join([]string{"BatchQueryJobReqPage", string(data)}, " ")
}
