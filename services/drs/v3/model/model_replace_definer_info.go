package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ReplaceDefinerInfo struct {
	JobId string `json:"job_id"`

	ReplaceDefiner bool `json:"replace_definer"`
}

func (o ReplaceDefinerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaceDefinerInfo struct{}"
	}

	return strings.Join([]string{"ReplaceDefinerInfo", string(data)}, " ")
}
