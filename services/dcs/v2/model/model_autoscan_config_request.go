package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AutoscanConfigRequest struct {
	EnableAutoScan *bool `json:"enable_auto_scan,omitempty"`

	ScheduleAt *[]string `json:"schedule_at,omitempty"`
}

func (o AutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"AutoscanConfigRequest", string(data)}, " ")
}
