package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlaveInstance struct {
	InstanceId string `json:"instance_id"`

	Region string `json:"region"`

	ProjectId string `json:"project_id"`

	ProjectName string `json:"project_name"`
}

func (o SlaveInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaveInstance struct{}"
	}

	return strings.Join([]string{"SlaveInstance", string(data)}, " ")
}
