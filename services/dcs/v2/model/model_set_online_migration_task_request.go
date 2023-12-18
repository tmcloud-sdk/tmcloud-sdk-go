package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetOnlineMigrationTaskRequest struct {
	TaskId string `json:"task_id"`

	Body *SetOnlineMigrationTaskBody `json:"body,omitempty"`
}

func (o SetOnlineMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOnlineMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"SetOnlineMigrationTaskRequest", string(data)}, " ")
}
