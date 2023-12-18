package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateScalingPolicyOption struct {
	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyType *UpdateScalingPolicyOptionScalingPolicyType `json:"scaling_policy_type,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV1 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`
}

func (o UpdateScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateScalingPolicyOption", string(data)}, " ")
}

type UpdateScalingPolicyOptionScalingPolicyType struct {
	value string
}

type UpdateScalingPolicyOptionScalingPolicyTypeEnum struct {
	ALARM      UpdateScalingPolicyOptionScalingPolicyType
	SCHEDULED  UpdateScalingPolicyOptionScalingPolicyType
	RECURRENCE UpdateScalingPolicyOptionScalingPolicyType
}

func GetUpdateScalingPolicyOptionScalingPolicyTypeEnum() UpdateScalingPolicyOptionScalingPolicyTypeEnum {
	return UpdateScalingPolicyOptionScalingPolicyTypeEnum{
		ALARM: UpdateScalingPolicyOptionScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: UpdateScalingPolicyOptionScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: UpdateScalingPolicyOptionScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c UpdateScalingPolicyOptionScalingPolicyType) Value() string {
	return c.value
}

func (c UpdateScalingPolicyOptionScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScalingPolicyOptionScalingPolicyType) UnmarshalJSON(b []byte) error {
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
