package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MasterInstance struct {
	InstanceId string `json:"instance_id"`

	Region string `json:"region"`

	ProjectId string `json:"project_id"`

	ProjectName string `json:"project_name"`
}

func (o MasterInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterInstance struct{}"
	}

	return strings.Join([]string{"MasterInstance", string(data)}, " ")
}
