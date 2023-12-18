package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListLoadBalancersRequest struct {
	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Description *[]string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`

	OperatingStatus *[]string `json:"operating_status,omitempty"`

	Guaranteed *bool `json:"guaranteed,omitempty"`

	VpcId *[]string `json:"vpc_id,omitempty"`

	VipPortId *[]string `json:"vip_port_id,omitempty"`

	VipAddress *[]string `json:"vip_address,omitempty"`

	VipSubnetCidrId *[]string `json:"vip_subnet_cidr_id,omitempty"`

	Ipv6VipPortId *[]string `json:"ipv6_vip_port_id,omitempty"`

	Ipv6VipAddress *[]string `json:"ipv6_vip_address,omitempty"`

	Ipv6VipVirsubnetId *[]string `json:"ipv6_vip_virsubnet_id,omitempty"`

	Eips *[]string `json:"eips,omitempty"`

	Publicips *[]string `json:"publicips,omitempty"`

	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`

	L4FlavorId *[]string `json:"l4_flavor_id,omitempty"`

	L4ScaleFlavorId *[]string `json:"l4_scale_flavor_id,omitempty"`

	L7FlavorId *[]string `json:"l7_flavor_id,omitempty"`

	L7ScaleFlavorId *[]string `json:"l7_scale_flavor_id,omitempty"`

	BillingInfo *[]string `json:"billing_info,omitempty"`

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	MemberAddress *[]string `json:"member_address,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	IpVersion *[]int32 `json:"ip_version,omitempty"`

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	ElbVirsubnetType *[]string `json:"elb_virsubnet_type,omitempty"`

	Autoscaling *[]string `json:"autoscaling,omitempty"`
}

func (o ListLoadBalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadBalancersRequest struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersRequest", string(data)}, " ")
}
