package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ScalingV1PolicyDetail struct {
	ScalingGroupId *string `json:"scaling_group_id,omitempty"`

	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	PolicyStatus *string `json:"policy_status,omitempty"`

	ScalingPolicyType *ScalingV1PolicyDetailScalingPolicyType `json:"scaling_policy_type,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV1 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`
}

func (o ScalingV1PolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingV1PolicyDetail struct{}"
	}

	return strings.Join([]string{"ScalingV1PolicyDetail", string(data)}, " ")
}

type ScalingV1PolicyDetailScalingPolicyType struct {
	value string
}

type ScalingV1PolicyDetailScalingPolicyTypeEnum struct {
	ALARM      ScalingV1PolicyDetailScalingPolicyType
	SCHEDULED  ScalingV1PolicyDetailScalingPolicyType
	RECURRENCE ScalingV1PolicyDetailScalingPolicyType
}

func GetScalingV1PolicyDetailScalingPolicyTypeEnum() ScalingV1PolicyDetailScalingPolicyTypeEnum {
	return ScalingV1PolicyDetailScalingPolicyTypeEnum{
		ALARM: ScalingV1PolicyDetailScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ScalingV1PolicyDetailScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ScalingV1PolicyDetailScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ScalingV1PolicyDetailScalingPolicyType) Value() string {
	return c.value
}

func (c ScalingV1PolicyDetailScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingV1PolicyDetailScalingPolicyType) UnmarshalJSON(b []byte) error {
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
