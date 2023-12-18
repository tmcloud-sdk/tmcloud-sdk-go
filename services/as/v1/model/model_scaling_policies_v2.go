package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingPoliciesV2 struct {
	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	ScalingResourceType *ScalingPoliciesV2ScalingResourceType `json:"scaling_resource_type,omitempty"`

	PolicyStatus *ScalingPoliciesV2PolicyStatus `json:"policy_status,omitempty"`

	ScalingPolicyType *ScalingPoliciesV2ScalingPolicyType `json:"scaling_policy_type,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV2 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	MetaData *ScalingPolicyV2MetaData `json:"meta_data,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ScalingPoliciesV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPoliciesV2 struct{}"
	}

	return strings.Join([]string{"ScalingPoliciesV2", string(data)}, " ")
}

type ScalingPoliciesV2ScalingResourceType struct {
	value string
}

type ScalingPoliciesV2ScalingResourceTypeEnum struct {
	SCALING_GROUP ScalingPoliciesV2ScalingResourceType
	BANDWIDTH     ScalingPoliciesV2ScalingResourceType
}

func GetScalingPoliciesV2ScalingResourceTypeEnum() ScalingPoliciesV2ScalingResourceTypeEnum {
	return ScalingPoliciesV2ScalingResourceTypeEnum{
		SCALING_GROUP: ScalingPoliciesV2ScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: ScalingPoliciesV2ScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c ScalingPoliciesV2ScalingResourceType) Value() string {
	return c.value
}

func (c ScalingPoliciesV2ScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPoliciesV2ScalingResourceType) UnmarshalJSON(b []byte) error {
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

type ScalingPoliciesV2PolicyStatus struct {
	value string
}

type ScalingPoliciesV2PolicyStatusEnum struct {
	INSERVICE ScalingPoliciesV2PolicyStatus
	PAUSED    ScalingPoliciesV2PolicyStatus
	EXECUTING ScalingPoliciesV2PolicyStatus
}

func GetScalingPoliciesV2PolicyStatusEnum() ScalingPoliciesV2PolicyStatusEnum {
	return ScalingPoliciesV2PolicyStatusEnum{
		INSERVICE: ScalingPoliciesV2PolicyStatus{
			value: "INSERVICE",
		},
		PAUSED: ScalingPoliciesV2PolicyStatus{
			value: "PAUSED",
		},
		EXECUTING: ScalingPoliciesV2PolicyStatus{
			value: "EXECUTING",
		},
	}
}

func (c ScalingPoliciesV2PolicyStatus) Value() string {
	return c.value
}

func (c ScalingPoliciesV2PolicyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPoliciesV2PolicyStatus) UnmarshalJSON(b []byte) error {
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

type ScalingPoliciesV2ScalingPolicyType struct {
	value string
}

type ScalingPoliciesV2ScalingPolicyTypeEnum struct {
	ALARM      ScalingPoliciesV2ScalingPolicyType
	SCHEDULED  ScalingPoliciesV2ScalingPolicyType
	RECURRENCE ScalingPoliciesV2ScalingPolicyType
}

func GetScalingPoliciesV2ScalingPolicyTypeEnum() ScalingPoliciesV2ScalingPolicyTypeEnum {
	return ScalingPoliciesV2ScalingPolicyTypeEnum{
		ALARM: ScalingPoliciesV2ScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ScalingPoliciesV2ScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ScalingPoliciesV2ScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ScalingPoliciesV2ScalingPolicyType) Value() string {
	return c.value
}

func (c ScalingPoliciesV2ScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPoliciesV2ScalingPolicyType) UnmarshalJSON(b []byte) error {
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
