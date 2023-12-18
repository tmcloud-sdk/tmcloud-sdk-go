package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodePool struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *NodePoolMetadata `json:"metadata"`

	Spec *NodePoolSpec `json:"spec"`

	Status *NodePoolStatus `json:"status,omitempty"`
}

func (o NodePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePool struct{}"
	}

	return strings.Join([]string{"NodePool", string(data)}, " ")
}
