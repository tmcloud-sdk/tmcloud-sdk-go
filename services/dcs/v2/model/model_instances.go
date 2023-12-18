package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Instances struct {
	InstanceId *string `json:"instance_id,omitempty"`

	InstanceName *string `json:"instance_name,omitempty"`
}

func (o Instances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instances struct{}"
	}

	return strings.Join([]string{"Instances", string(data)}, " ")
}
