package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateMembersOption struct {
	Name *string `json:"name,omitempty"`

	Address string `json:"address"`

	ProtocolPort int32 `json:"protocol_port"`

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchCreateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersOption", string(data)}, " ")
}
