package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateDatabaseObjectReq struct {
	JobId string `json:"job_id"`

	Selected *bool `json:"selected,omitempty"`

	SyncDatabase *bool `json:"sync_database,omitempty"`

	Job *[]DatabaseInfo `json:"job,omitempty"`
}

func (o UpdateDatabaseObjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseObjectReq struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseObjectReq", string(data)}, " ")
}
