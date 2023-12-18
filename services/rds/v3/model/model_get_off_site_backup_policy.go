package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GetOffSiteBackupPolicy struct {
	BackupType *string `json:"backup_type,omitempty"`

	KeepDays *int32 `json:"keep_days,omitempty"`

	DestinationRegion *string `json:"destination_region,omitempty"`

	DestinationProjectId *string `json:"destination_project_id,omitempty"`
}

func (o GetOffSiteBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetOffSiteBackupPolicy struct{}"
	}

	return strings.Join([]string{"GetOffSiteBackupPolicy", string(data)}, " ")
}
