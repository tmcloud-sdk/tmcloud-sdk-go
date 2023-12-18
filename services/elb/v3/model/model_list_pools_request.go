package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListPoolsRequest struct {
	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Description *[]string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	HealthmonitorId *[]string `json:"healthmonitor_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	Protocol *[]string `json:"protocol,omitempty"`

	LbAlgorithm *[]string `json:"lb_algorithm,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	IpVersion *[]string `json:"ip_version,omitempty"`

	MemberAddress *[]string `json:"member_address,omitempty"`

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	ListenerId *[]string `json:"listener_id,omitempty"`

	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`

	VpcId *[]string `json:"vpc_id,omitempty"`

	Type *[]string `json:"type,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
