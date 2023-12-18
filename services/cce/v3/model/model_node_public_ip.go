package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodePublicIp struct {
	Ids *[]string `json:"ids,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Eip *NodeEipSpec `json:"eip,omitempty"`
}

func (o NodePublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePublicIp struct{}"
	}

	return strings.Join([]string{"NodePublicIp", string(data)}, " ")
}
