package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BackupStrategy struct {
	StartTime string `json:"start_time"`

	KeepDays *int32 `json:"keep_days,omitempty"`
}

func (o BackupStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategy struct{}"
	}

	return strings.Join([]string{"BackupStrategy", string(data)}, " ")
}
