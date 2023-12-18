package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceNodesInfoResp struct {
	InstanceId *string `json:"instance_id,omitempty"`

	NodeCount *int32 `json:"node_count,omitempty"`

	Nodes *[]NodesInfoResp `json:"nodes,omitempty"`
}

func (o InstanceNodesInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceNodesInfoResp struct{}"
	}

	return strings.Join([]string{"InstanceNodesInfoResp", string(data)}, " ")
}
