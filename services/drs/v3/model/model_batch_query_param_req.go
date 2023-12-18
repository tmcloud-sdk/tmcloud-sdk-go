package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchQueryParamReq struct {
	Jobs []string `json:"jobs"`

	Refresh string `json:"refresh"`
}

func (o BatchQueryParamReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryParamReq struct{}"
	}

	return strings.Join([]string{"BatchQueryParamReq", string(data)}, " ")
}
