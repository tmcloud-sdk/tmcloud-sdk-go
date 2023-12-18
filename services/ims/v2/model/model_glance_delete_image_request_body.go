package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceDeleteImageRequestBody struct {
	DeleteBackup *bool `json:"delete_backup,omitempty"`
}

func (o GlanceDeleteImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageRequestBody", string(data)}, " ")
}
