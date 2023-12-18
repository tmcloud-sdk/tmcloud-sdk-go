package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteMigrateTaskRequest struct {
	TaskIdList []string `json:"task_id_list"`
}

func (o DeleteMigrateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrateTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrateTaskRequest", string(data)}, " ")
}
