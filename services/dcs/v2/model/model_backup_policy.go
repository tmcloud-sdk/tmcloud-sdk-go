package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BackupPolicy struct {
	BackupType string `json:"backup_type"`

	SaveDays *int32 `json:"save_days,omitempty"`

	PeriodicalBackupPlan *BackupPlan `json:"periodical_backup_plan,omitempty"`
}

func (o BackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicy struct{}"
	}

	return strings.Join([]string{"BackupPolicy", string(data)}, " ")
}
