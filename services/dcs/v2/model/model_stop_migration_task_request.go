package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StopMigrationTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o StopMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskRequest", string(data)}, " ")
}
