package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSwitchoverReq struct {
	Jobs []string `json:"jobs"`
}

func (o BatchSwitchoverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchoverReq struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverReq", string(data)}, " ")
}
