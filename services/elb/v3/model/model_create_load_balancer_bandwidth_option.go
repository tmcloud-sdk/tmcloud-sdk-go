package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateLoadBalancerBandwidthOption struct {
	Name *string `json:"name,omitempty"`

	Size *int32 `json:"size,omitempty"`

	ChargeMode *CreateLoadBalancerBandwidthOptionChargeMode `json:"charge_mode,omitempty"`

	ShareType *CreateLoadBalancerBandwidthOptionShareType `json:"share_type,omitempty"`

	BillingInfo *string `json:"billing_info,omitempty"`

	Id *string `json:"id,omitempty"`
}

func (o CreateLoadBalancerBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerBandwidthOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerBandwidthOption", string(data)}, " ")
}

type CreateLoadBalancerBandwidthOptionChargeMode struct {
	value string
}

type CreateLoadBalancerBandwidthOptionChargeModeEnum struct {
	BANDWIDTH CreateLoadBalancerBandwidthOptionChargeMode
	TRAFFIC   CreateLoadBalancerBandwidthOptionChargeMode
}

func GetCreateLoadBalancerBandwidthOptionChargeModeEnum() CreateLoadBalancerBandwidthOptionChargeModeEnum {
	return CreateLoadBalancerBandwidthOptionChargeModeEnum{
		BANDWIDTH: CreateLoadBalancerBandwidthOptionChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: CreateLoadBalancerBandwidthOptionChargeMode{
			value: "traffic",
		},
	}
}

func (c CreateLoadBalancerBandwidthOptionChargeMode) Value() string {
	return c.value
}

func (c CreateLoadBalancerBandwidthOptionChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerBandwidthOptionChargeMode) UnmarshalJSON(b []byte) error {
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

type CreateLoadBalancerBandwidthOptionShareType struct {
	value string
}

type CreateLoadBalancerBandwidthOptionShareTypeEnum struct {
	PER   CreateLoadBalancerBandwidthOptionShareType
	WHOLE CreateLoadBalancerBandwidthOptionShareType
}

func GetCreateLoadBalancerBandwidthOptionShareTypeEnum() CreateLoadBalancerBandwidthOptionShareTypeEnum {
	return CreateLoadBalancerBandwidthOptionShareTypeEnum{
		PER: CreateLoadBalancerBandwidthOptionShareType{
			value: "PER",
		},
		WHOLE: CreateLoadBalancerBandwidthOptionShareType{
			value: "WHOLE",
		},
	}
}

func (c CreateLoadBalancerBandwidthOptionShareType) Value() string {
	return c.value
}

func (c CreateLoadBalancerBandwidthOptionShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerBandwidthOptionShareType) UnmarshalJSON(b []byte) error {
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
