package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateHealthMonitorResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Healthmonitor  *HealthMonitor `json:"healthmonitor,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateHealthMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorResponse", string(data)}, " ")
}
