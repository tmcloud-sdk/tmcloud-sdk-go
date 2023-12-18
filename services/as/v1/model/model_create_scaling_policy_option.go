package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateScalingPolicyOption struct {
	ScalingPolicyName string `json:"scaling_policy_name"`

	ScalingGroupId string `json:"scaling_group_id"`

	ScalingPolicyType CreateScalingPolicyOptionScalingPolicyType `json:"scaling_policy_type"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV1 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`
}

func (o CreateScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyOption", string(data)}, " ")
}

type CreateScalingPolicyOptionScalingPolicyType struct {
	value string
}

type CreateScalingPolicyOptionScalingPolicyTypeEnum struct {
	ALARM      CreateScalingPolicyOptionScalingPolicyType
	SCHEDULED  CreateScalingPolicyOptionScalingPolicyType
	RECURRENCE CreateScalingPolicyOptionScalingPolicyType
}

func GetCreateScalingPolicyOptionScalingPolicyTypeEnum() CreateScalingPolicyOptionScalingPolicyTypeEnum {
	return CreateScalingPolicyOptionScalingPolicyTypeEnum{
		ALARM: CreateScalingPolicyOptionScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: CreateScalingPolicyOptionScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: CreateScalingPolicyOptionScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c CreateScalingPolicyOptionScalingPolicyType) Value() string {
	return c.value
}

func (c CreateScalingPolicyOptionScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingPolicyOptionScalingPolicyType) UnmarshalJSON(b []byte) error {
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
