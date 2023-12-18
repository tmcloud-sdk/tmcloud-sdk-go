package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateManualBackupRequestBody struct {
	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Databases *[]BackupDatabase `json:"databases,omitempty"`
}

func (o CreateManualBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateManualBackupRequestBody", string(data)}, " ")
}
