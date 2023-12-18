package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StopMigrationTaskSyncRequest struct {
	TaskId string `json:"task_id"`
}

func (o StopMigrationTaskSyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskSyncRequest struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskSyncRequest", string(data)}, " ")
}
