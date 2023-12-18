package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateScalingPolicyV2Option struct {
	ScalingPolicyName string `json:"scaling_policy_name"`

	ScalingResourceId string `json:"scaling_resource_id"`

	ScalingResourceType CreateScalingPolicyV2OptionScalingResourceType `json:"scaling_resource_type"`

	ScalingPolicyType CreateScalingPolicyV2OptionScalingPolicyType `json:"scaling_policy_type"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV2 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o CreateScalingPolicyV2Option) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyV2Option struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyV2Option", string(data)}, " ")
}

type CreateScalingPolicyV2OptionScalingResourceType struct {
	value string
}

type CreateScalingPolicyV2OptionScalingResourceTypeEnum struct {
	SCALING_GROUP CreateScalingPolicyV2OptionScalingResourceType
	BANDWIDTH     CreateScalingPolicyV2OptionScalingResourceType
}

func GetCreateScalingPolicyV2OptionScalingResourceTypeEnum() CreateScalingPolicyV2OptionScalingResourceTypeEnum {
	return CreateScalingPolicyV2OptionScalingResourceTypeEnum{
		SCALING_GROUP: CreateScalingPolicyV2OptionScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: CreateScalingPolicyV2OptionScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c CreateScalingPolicyV2OptionScalingResourceType) Value() string {
	return c.value
}

func (c CreateScalingPolicyV2OptionScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingPolicyV2OptionScalingResourceType) UnmarshalJSON(b []byte) error {
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

type CreateScalingPolicyV2OptionScalingPolicyType struct {
	value string
}

type CreateScalingPolicyV2OptionScalingPolicyTypeEnum struct {
	ALARM      CreateScalingPolicyV2OptionScalingPolicyType
	SCHEDULED  CreateScalingPolicyV2OptionScalingPolicyType
	RECURRENCE CreateScalingPolicyV2OptionScalingPolicyType
}

func GetCreateScalingPolicyV2OptionScalingPolicyTypeEnum() CreateScalingPolicyV2OptionScalingPolicyTypeEnum {
	return CreateScalingPolicyV2OptionScalingPolicyTypeEnum{
		ALARM: CreateScalingPolicyV2OptionScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: CreateScalingPolicyV2OptionScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: CreateScalingPolicyV2OptionScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c CreateScalingPolicyV2OptionScalingPolicyType) Value() string {
	return c.value
}

func (c CreateScalingPolicyV2OptionScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingPolicyV2OptionScalingPolicyType) UnmarshalJSON(b []byte) error {
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
