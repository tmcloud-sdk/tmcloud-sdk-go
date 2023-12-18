package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceDrRelation struct {
	InstanceId *string `json:"instance_id,omitempty"`

	MasterInstance *MasterInstance `json:"master_instance,omitempty"`

	SlaveInstances *[]SlaveInstance `json:"slave_instances,omitempty"`
}

func (o InstanceDrRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDrRelation struct{}"
	}

	return strings.Join([]string{"InstanceDrRelation", string(data)}, " ")
}
