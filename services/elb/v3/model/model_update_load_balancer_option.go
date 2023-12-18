package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateLoadBalancerOption struct {
	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	VipAddress *string `json:"vip_address,omitempty"`

	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidUpdateOption `json:"prepaid_options,omitempty"`

	Autoscaling *UpdateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`

	WafFailureAction *UpdateLoadBalancerOptionWafFailureAction `json:"waf_failure_action,omitempty"`
}

func (o UpdateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerOption", string(data)}, " ")
}

type UpdateLoadBalancerOptionWafFailureAction struct {
	value string
}

type UpdateLoadBalancerOptionWafFailureActionEnum struct {
	DISCARD UpdateLoadBalancerOptionWafFailureAction
	FORWARD UpdateLoadBalancerOptionWafFailureAction
}

func GetUpdateLoadBalancerOptionWafFailureActionEnum() UpdateLoadBalancerOptionWafFailureActionEnum {
	return UpdateLoadBalancerOptionWafFailureActionEnum{
		DISCARD: UpdateLoadBalancerOptionWafFailureAction{
			value: "discard",
		},
		FORWARD: UpdateLoadBalancerOptionWafFailureAction{
			value: "forward",
		},
	}
}

func (c UpdateLoadBalancerOptionWafFailureAction) Value() string {
	return c.value
}

func (c UpdateLoadBalancerOptionWafFailureAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLoadBalancerOptionWafFailureAction) UnmarshalJSON(b []byte) error {
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
