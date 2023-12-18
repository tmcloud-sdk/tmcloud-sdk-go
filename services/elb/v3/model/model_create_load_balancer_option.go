package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateLoadBalancerOption struct {
	Id *string `json:"id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	VipAddress *string `json:"vip_address,omitempty"`

	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	Provider *string `json:"provider,omitempty"`

	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	Guaranteed *bool `json:"guaranteed,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	AvailabilityZoneList []string `json:"availability_zone_list"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	BillingInfo *string `json:"billing_info,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	Publicip *CreateLoadBalancerPublicIpOption `json:"publicip,omitempty"`

	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidCreateOption `json:"prepaid_options,omitempty"`

	Autoscaling *CreateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`

	WafFailureAction *CreateLoadBalancerOptionWafFailureAction `json:"waf_failure_action,omitempty"`
}

func (o CreateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerOption", string(data)}, " ")
}

type CreateLoadBalancerOptionWafFailureAction struct {
	value string
}

type CreateLoadBalancerOptionWafFailureActionEnum struct {
	DISCARD CreateLoadBalancerOptionWafFailureAction
	FORWARD CreateLoadBalancerOptionWafFailureAction
}

func GetCreateLoadBalancerOptionWafFailureActionEnum() CreateLoadBalancerOptionWafFailureActionEnum {
	return CreateLoadBalancerOptionWafFailureActionEnum{
		DISCARD: CreateLoadBalancerOptionWafFailureAction{
			value: "discard",
		},
		FORWARD: CreateLoadBalancerOptionWafFailureAction{
			value: "forward",
		},
	}
}

func (c CreateLoadBalancerOptionWafFailureAction) Value() string {
	return c.value
}

func (c CreateLoadBalancerOptionWafFailureAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerOptionWafFailureAction) UnmarshalJSON(b []byte) error {
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
