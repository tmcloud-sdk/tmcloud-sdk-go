package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchQueryProgressReq struct {
	Jobs []string `json:"jobs"`
}

func (o BatchQueryProgressReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryProgressReq struct{}"
	}

	return strings.Join([]string{"BatchQueryProgressReq", string(data)}, " ")
}
