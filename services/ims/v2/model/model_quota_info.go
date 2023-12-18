package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QuotaInfo struct {
	Type string `json:"type"`

	Used int32 `json:"used"`

	Quota int32 `json:"quota"`

	Min int32 `json:"min"`

	Max int32 `json:"max"`
}

func (o QuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInfo struct{}"
	}

	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
