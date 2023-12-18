package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QuotaDetailVolumes struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailVolumes struct{}"
	}

	return strings.Join([]string{"QuotaDetailVolumes", string(data)}, " ")
}
