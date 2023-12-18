package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteHotkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`

	HotkeyId string `json:"hotkey_id"`
}

func (o DeleteHotkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteHotkeyScanTaskRequest", string(data)}, " ")
}
