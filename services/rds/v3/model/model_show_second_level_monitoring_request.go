package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowSecondLevelMonitoringRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowSecondLevelMonitoringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecondLevelMonitoringRequest struct{}"
	}

	return strings.Join([]string{"ShowSecondLevelMonitoringRequest", string(data)}, " ")
}
