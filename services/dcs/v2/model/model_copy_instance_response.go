package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CopyInstanceResponse struct {
	BackupId       *string `json:"backup_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceResponse struct{}"
	}

	return strings.Join([]string{"CopyInstanceResponse", string(data)}, " ")
}
