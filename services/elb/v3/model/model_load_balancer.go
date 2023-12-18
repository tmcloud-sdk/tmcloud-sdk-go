package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type LoadBalancer struct {
	Id string `json:"id"`

	Description string `json:"description"`

	ProvisioningStatus string `json:"provisioning_status"`

	AdminStateUp bool `json:"admin_state_up"`

	Provider string `json:"provider"`

	Pools []PoolRef `json:"pools"`

	Listeners []ListenerRef `json:"listeners"`

	OperatingStatus string `json:"operating_status"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	VipSubnetCidrId string `json:"vip_subnet_cidr_id"`

	VipAddress string `json:"vip_address"`

	VipPortId string `json:"vip_port_id"`

	Tags []Tag `json:"tags"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`

	Guaranteed bool `json:"guaranteed"`

	VpcId string `json:"vpc_id"`

	Eips []EipInfo `json:"eips"`

	Ipv6VipAddress string `json:"ipv6_vip_address"`

	Ipv6VipVirsubnetId string `json:"ipv6_vip_virsubnet_id"`

	Ipv6VipPortId string `json:"ipv6_vip_port_id"`

	AvailabilityZoneList []string `json:"availability_zone_list"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	BillingInfo string `json:"billing_info"`

	L4FlavorId string `json:"l4_flavor_id"`

	L4ScaleFlavorId string `json:"l4_scale_flavor_id"`

	L7FlavorId string `json:"l7_flavor_id"`

	L7ScaleFlavorId string `json:"l7_scale_flavor_id"`

	Publicips []PublicIpInfo `json:"publicips"`

	GlobalEips []GlobalEipInfo `json:"global_eips"`

	ElbVirsubnetIds []string `json:"elb_virsubnet_ids"`

	ElbVirsubnetType LoadBalancerElbVirsubnetType `json:"elb_virsubnet_type"`

	IpTargetEnable bool `json:"ip_target_enable"`

	FrozenScene string `json:"frozen_scene"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth"`

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	Autoscaling *AutoscalingRef `json:"autoscaling,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	WafFailureAction *string `json:"waf_failure_action,omitempty"`
}

func (o LoadBalancer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancer struct{}"
	}

	return strings.Join([]string{"LoadBalancer", string(data)}, " ")
}

type LoadBalancerElbVirsubnetType struct {
	value string
}

type LoadBalancerElbVirsubnetTypeEnum struct {
	IPV4      LoadBalancerElbVirsubnetType
	DUALSTACK LoadBalancerElbVirsubnetType
}

func GetLoadBalancerElbVirsubnetTypeEnum() LoadBalancerElbVirsubnetTypeEnum {
	return LoadBalancerElbVirsubnetTypeEnum{
		IPV4: LoadBalancerElbVirsubnetType{
			value: "ipv4",
		},
		DUALSTACK: LoadBalancerElbVirsubnetType{
			value: "dualstack",
		},
	}
}

func (c LoadBalancerElbVirsubnetType) Value() string {
	return c.value
}

func (c LoadBalancerElbVirsubnetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadBalancerElbVirsubnetType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
