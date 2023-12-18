package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowHotkeyAutoscanConfigResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	EnableAutoScan *bool `json:"enable_auto_scan,omitempty"`

	ScheduleAt *[]string `json:"schedule_at,omitempty"`

	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHotkeyAutoscanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyAutoscanConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowHotkeyAutoscanConfigResponse", string(data)}, " ")
}
