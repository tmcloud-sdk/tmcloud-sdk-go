package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingAllPolicyDetail struct {
	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	ScalingResourceType *ScalingAllPolicyDetailScalingResourceType `json:"scaling_resource_type,omitempty"`

	PolicyStatus *ScalingAllPolicyDetailPolicyStatus `json:"policy_status,omitempty"`

	ScalingPolicyType *ScalingAllPolicyDetailScalingPolicyType `json:"scaling_policy_type,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV2 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	MetaData *ScalingPolicyV2MetaData `json:"meta_data,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ScalingAllPolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingAllPolicyDetail struct{}"
	}

	return strings.Join([]string{"ScalingAllPolicyDetail", string(data)}, " ")
}

type ScalingAllPolicyDetailScalingResourceType struct {
	value string
}

type ScalingAllPolicyDetailScalingResourceTypeEnum struct {
	SCALING_GROUP ScalingAllPolicyDetailScalingResourceType
	BANDWIDTH     ScalingAllPolicyDetailScalingResourceType
}

func GetScalingAllPolicyDetailScalingResourceTypeEnum() ScalingAllPolicyDetailScalingResourceTypeEnum {
	return ScalingAllPolicyDetailScalingResourceTypeEnum{
		SCALING_GROUP: ScalingAllPolicyDetailScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: ScalingAllPolicyDetailScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c ScalingAllPolicyDetailScalingResourceType) Value() string {
	return c.value
}

func (c ScalingAllPolicyDetailScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingAllPolicyDetailScalingResourceType) UnmarshalJSON(b []byte) error {
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

type ScalingAllPolicyDetailPolicyStatus struct {
	value string
}

type ScalingAllPolicyDetailPolicyStatusEnum struct {
	INSERVICE ScalingAllPolicyDetailPolicyStatus
	PAUSED    ScalingAllPolicyDetailPolicyStatus
	EXECUTING ScalingAllPolicyDetailPolicyStatus
}

func GetScalingAllPolicyDetailPolicyStatusEnum() ScalingAllPolicyDetailPolicyStatusEnum {
	return ScalingAllPolicyDetailPolicyStatusEnum{
		INSERVICE: ScalingAllPolicyDetailPolicyStatus{
			value: "INSERVICE",
		},
		PAUSED: ScalingAllPolicyDetailPolicyStatus{
			value: "PAUSED",
		},
		EXECUTING: ScalingAllPolicyDetailPolicyStatus{
			value: "EXECUTING",
		},
	}
}

func (c ScalingAllPolicyDetailPolicyStatus) Value() string {
	return c.value
}

func (c ScalingAllPolicyDetailPolicyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingAllPolicyDetailPolicyStatus) UnmarshalJSON(b []byte) error {
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

type ScalingAllPolicyDetailScalingPolicyType struct {
	value string
}

type ScalingAllPolicyDetailScalingPolicyTypeEnum struct {
	ALARM      ScalingAllPolicyDetailScalingPolicyType
	SCHEDULED  ScalingAllPolicyDetailScalingPolicyType
	RECURRENCE ScalingAllPolicyDetailScalingPolicyType
}

func GetScalingAllPolicyDetailScalingPolicyTypeEnum() ScalingAllPolicyDetailScalingPolicyTypeEnum {
	return ScalingAllPolicyDetailScalingPolicyTypeEnum{
		ALARM: ScalingAllPolicyDetailScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ScalingAllPolicyDetailScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ScalingAllPolicyDetailScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ScalingAllPolicyDetailScalingPolicyType) Value() string {
	return c.value
}

func (c ScalingAllPolicyDetailScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingAllPolicyDetailScalingPolicyType) UnmarshalJSON(b []byte) error {
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
