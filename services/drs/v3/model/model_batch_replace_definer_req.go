package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchReplaceDefinerReq struct {
	Jobs []ReplaceDefinerInfo `json:"jobs"`
}

func (o BatchReplaceDefinerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchReplaceDefinerReq struct{}"
	}

	return strings.Join([]string{"BatchReplaceDefinerReq", string(data)}, " ")
}
