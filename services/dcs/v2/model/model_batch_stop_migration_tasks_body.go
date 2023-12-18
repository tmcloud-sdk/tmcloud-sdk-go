package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchStopMigrationTasksBody struct {
	MigrationTasks []string `json:"migration_tasks"`
}

func (o BatchStopMigrationTasksBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopMigrationTasksBody struct{}"
	}

	return strings.Join([]string{"BatchStopMigrationTasksBody", string(data)}, " ")
}
