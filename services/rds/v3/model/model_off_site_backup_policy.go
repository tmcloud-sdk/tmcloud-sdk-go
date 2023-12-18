package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type OffSiteBackupPolicy struct {
	BackupType string `json:"backup_type"`

	KeepDays int32 `json:"keep_days"`

	DestinationRegion string `json:"destination_region"`

	DestinationProjectId string `json:"destination_project_id"`
}

func (o OffSiteBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OffSiteBackupPolicy struct{}"
	}

	return strings.Join([]string{"OffSiteBackupPolicy", string(data)}, " ")
}
