package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMembersRequest struct {
	PoolId string `json:"pool_id"`

	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Weight *[]int32 `json:"weight,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	SubnetCidrId *[]string `json:"subnet_cidr_id,omitempty"`

	Address *[]string `json:"address,omitempty"`

	ProtocolPort *[]int32 `json:"protocol_port,omitempty"`

	Id *[]string `json:"id,omitempty"`

	OperatingStatus *[]string `json:"operating_status,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	IpVersion *[]string `json:"ip_version,omitempty"`

	MemberType *[]string `json:"member_type,omitempty"`

	InstanceId *[]string `json:"instance_id,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
