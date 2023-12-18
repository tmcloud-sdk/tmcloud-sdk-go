package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StopMigrationTaskSyncResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopMigrationTaskSyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskSyncResponse struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskSyncResponse", string(data)}, " ")
}
