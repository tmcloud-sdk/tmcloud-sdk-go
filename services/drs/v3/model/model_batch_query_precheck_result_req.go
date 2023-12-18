package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchQueryPrecheckResultReq struct {
	Jobs []string `json:"jobs"`
}

func (o BatchQueryPrecheckResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryPrecheckResultReq struct{}"
	}

	return strings.Join([]string{"BatchQueryPrecheckResultReq", string(data)}, " ")
}
