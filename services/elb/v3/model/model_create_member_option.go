package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateMemberOption struct {
	Address string `json:"address"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	ProtocolPort int32 `json:"protocol_port"`

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o CreateMemberOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberOption struct{}"
	}

	return strings.Join([]string{"CreateMemberOption", string(data)}, " ")
}
