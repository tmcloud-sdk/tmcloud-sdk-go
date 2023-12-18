package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Member struct {
	Id string `json:"id"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	PoolId *string `json:"pool_id,omitempty"`

	AdminStateUp bool `json:"admin_state_up"`

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	ProtocolPort int32 `json:"protocol_port"`

	Weight int32 `json:"weight"`

	Address string `json:"address"`

	IpVersion string `json:"ip_version"`

	DeviceOwner *string `json:"device_owner,omitempty"`

	DeviceId *string `json:"device_id,omitempty"`

	OperatingStatus string `json:"operating_status"`

	Status []MemberStatus `json:"status"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	Loadbalancers *[]ResourceId `json:"loadbalancers,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	MemberType *string `json:"member_type,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
