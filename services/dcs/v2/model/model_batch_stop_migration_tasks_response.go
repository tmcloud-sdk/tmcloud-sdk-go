package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchStopMigrationTasksResponse struct {
	MigrationTasks *[]StopMigrationTaskResult `json:"migration_tasks,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchStopMigrationTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopMigrationTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchStopMigrationTasksResponse", string(data)}, " ")
}
