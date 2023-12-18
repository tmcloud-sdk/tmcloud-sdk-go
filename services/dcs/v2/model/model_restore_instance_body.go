package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RestoreInstanceBody struct {
	BackupId string `json:"backup_id"`

	Remark *string `json:"remark,omitempty"`
}

func (o RestoreInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceBody", string(data)}, " ")
}
