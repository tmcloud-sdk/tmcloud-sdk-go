package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowHotkeyTaskDetailsRequest struct {
	InstanceId string `json:"instance_id"`

	HotkeyId string `json:"hotkey_id"`
}

func (o ShowHotkeyTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowHotkeyTaskDetailsRequest", string(data)}, " ")
}
