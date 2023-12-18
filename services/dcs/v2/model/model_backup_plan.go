package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BackupPlan struct {
	TimezoneOffset *string `json:"timezone_offset,omitempty"`

	BackupAt []int32 `json:"backup_at"`

	PeriodType string `json:"period_type"`

	BeginAt string `json:"begin_at"`
}

func (o BackupPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPlan struct{}"
	}

	return strings.Join([]string{"BackupPlan", string(data)}, " ")
}
