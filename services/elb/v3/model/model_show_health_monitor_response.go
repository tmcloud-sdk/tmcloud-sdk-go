package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowHealthMonitorResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Healthmonitor  *HealthMonitor `json:"healthmonitor,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHealthMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthMonitorResponse", string(data)}, " ")
}
