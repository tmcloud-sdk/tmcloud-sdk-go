package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BackupStrategyForResponse struct {
	StartTime string `json:"start_time"`

	KeepDays int32 `json:"keep_days"`
}

func (o BackupStrategyForResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyForResponse struct{}"
	}

	return strings.Join([]string{"BackupStrategyForResponse", string(data)}, " ")
}
