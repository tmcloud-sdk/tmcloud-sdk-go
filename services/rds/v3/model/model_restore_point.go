package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RestorePoint struct {
	InstanceId string `json:"instance_id"`

	Type string `json:"type"`

	BackupId *string `json:"backup_id,omitempty"`

	RestoreTime *int64 `json:"restore_time,omitempty"`

	DatabaseName map[string]string `json:"database_name,omitempty"`
}

func (o RestorePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePoint struct{}"
	}

	return strings.Join([]string{"RestorePoint", string(data)}, " ")
}
