package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListBackupRecordsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	BackupRecordResponse *[]BackupRecordResponse `json:"backup_record_response,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListBackupRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupRecordsResponse", string(data)}, " ")
}
