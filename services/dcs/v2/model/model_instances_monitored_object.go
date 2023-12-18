package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstancesMonitoredObject struct {
	DcsInstanceId *string `json:"dcs_instance_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o InstancesMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesMonitoredObject struct{}"
	}

	return strings.Join([]string{"InstancesMonitoredObject", string(data)}, " ")
}
