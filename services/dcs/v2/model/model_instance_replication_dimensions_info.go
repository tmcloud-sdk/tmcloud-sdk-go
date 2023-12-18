package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceReplicationDimensionsInfo struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o InstanceReplicationDimensionsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceReplicationDimensionsInfo struct{}"
	}

	return strings.Join([]string{"InstanceReplicationDimensionsInfo", string(data)}, " ")
}
