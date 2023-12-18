package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RunlogItem struct {
	Id *string `json:"id,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	ReplicationIp *string `json:"replication_ip,omitempty"`

	Status *string `json:"status,omitempty"`

	Time *string `json:"time,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`
}

func (o RunlogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunlogItem struct{}"
	}

	return strings.Join([]string{"RunlogItem", string(data)}, " ")
}
