package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateScalingV2PolicyOption struct {
	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	ScalingResourceType *UpdateScalingV2PolicyOptionScalingResourceType `json:"scaling_resource_type,omitempty"`

	ScalingPolicyType *UpdateScalingV2PolicyOptionScalingPolicyType `json:"scaling_policy_type,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV2 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateScalingV2PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingV2PolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateScalingV2PolicyOption", string(data)}, " ")
}

type UpdateScalingV2PolicyOptionScalingResourceType struct {
	value string
}

type UpdateScalingV2PolicyOptionScalingResourceTypeEnum struct {
	SCALING_GROUP UpdateScalingV2PolicyOptionScalingResourceType
	BANDWIDTH     UpdateScalingV2PolicyOptionScalingResourceType
}

func GetUpdateScalingV2PolicyOptionScalingResourceTypeEnum() UpdateScalingV2PolicyOptionScalingResourceTypeEnum {
	return UpdateScalingV2PolicyOptionScalingResourceTypeEnum{
		SCALING_GROUP: UpdateScalingV2PolicyOptionScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: UpdateScalingV2PolicyOptionScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c UpdateScalingV2PolicyOptionScalingResourceType) Value() string {
	return c.value
}

func (c UpdateScalingV2PolicyOptionScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScalingV2PolicyOptionScalingResourceType) UnmarshalJSON(b []byte) error {
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

type UpdateScalingV2PolicyOptionScalingPolicyType struct {
	value string
}

type UpdateScalingV2PolicyOptionScalingPolicyTypeEnum struct {
	ALARM      UpdateScalingV2PolicyOptionScalingPolicyType
	SCHEDULED  UpdateScalingV2PolicyOptionScalingPolicyType
	RECURRENCE UpdateScalingV2PolicyOptionScalingPolicyType
}

func GetUpdateScalingV2PolicyOptionScalingPolicyTypeEnum() UpdateScalingV2PolicyOptionScalingPolicyTypeEnum {
	return UpdateScalingV2PolicyOptionScalingPolicyTypeEnum{
		ALARM: UpdateScalingV2PolicyOptionScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: UpdateScalingV2PolicyOptionScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: UpdateScalingV2PolicyOptionScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c UpdateScalingV2PolicyOptionScalingPolicyType) Value() string {
	return c.value
}

func (c UpdateScalingV2PolicyOptionScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScalingV2PolicyOptionScalingPolicyType) UnmarshalJSON(b []byte) error {
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
