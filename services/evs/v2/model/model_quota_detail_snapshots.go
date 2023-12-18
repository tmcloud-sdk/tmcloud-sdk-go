package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QuotaDetailSnapshots struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailSnapshots) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailSnapshots struct{}"
	}

	return strings.Join([]string{"QuotaDetailSnapshots", string(data)}, " ")
}
