package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchMember struct {
	Id string `json:"id"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	AdminStateUp bool `json:"admin_state_up"`

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	ProtocolPort int32 `json:"protocol_port"`

	Weight int32 `json:"weight"`

	Address string `json:"address"`

	OperatingStatus string `json:"operating_status"`

	Status *[]MemberStatus `json:"status,omitempty"`

	MemberType *string `json:"member_type,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	PortId string `json:"port_id"`

	RetStatus string `json:"ret_status"`
}

func (o BatchMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMember struct{}"
	}

	return strings.Join([]string{"BatchMember", string(data)}, " ")
}
