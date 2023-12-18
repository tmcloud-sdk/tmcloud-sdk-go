package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetSecondLevelMonitorRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`

	Body *SecondMonitor `json:"body,omitempty"`
}

func (o SetSecondLevelMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSecondLevelMonitorRequest struct{}"
	}

	return strings.Join([]string{"SetSecondLevelMonitorRequest", string(data)}, " ")
}
