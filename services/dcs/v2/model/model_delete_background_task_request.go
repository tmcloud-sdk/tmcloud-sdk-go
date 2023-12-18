package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteBackgroundTaskRequest struct {
	InstanceId string `json:"instance_id"`

	TaskId string `json:"task_id"`
}

func (o DeleteBackgroundTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackgroundTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackgroundTaskRequest", string(data)}, " ")
}
