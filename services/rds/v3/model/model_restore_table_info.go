package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RestoreTableInfo struct {
	OldName string `json:"oldName"`

	NewName string `json:"newName"`
}

func (o RestoreTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTableInfo struct{}"
	}

	return strings.Join([]string{"RestoreTableInfo", string(data)}, " ")
}
