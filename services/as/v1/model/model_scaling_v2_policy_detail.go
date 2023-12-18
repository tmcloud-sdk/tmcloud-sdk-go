package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingV2PolicyDetail struct {
	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	ScalingResourceType *ScalingV2PolicyDetailScalingResourceType `json:"scaling_resource_type,omitempty"`

	PolicyStatus *ScalingV2PolicyDetailPolicyStatus `json:"policy_status,omitempty"`

	ScalingPolicyType *ScalingV2PolicyDetailScalingPolicyType `json:"scaling_policy_type,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV2 `json:"scaling_policy_action,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	MetaData *ScalingPolicyV2MetaData `json:"meta_data,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ScalingV2PolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingV2PolicyDetail struct{}"
	}

	return strings.Join([]string{"ScalingV2PolicyDetail", string(data)}, " ")
}

type ScalingV2PolicyDetailScalingResourceType struct {
	value string
}

type ScalingV2PolicyDetailScalingResourceTypeEnum struct {
	SCALING_GROUP ScalingV2PolicyDetailScalingResourceType
	BANDWIDTH     ScalingV2PolicyDetailScalingResourceType
}

func GetScalingV2PolicyDetailScalingResourceTypeEnum() ScalingV2PolicyDetailScalingResourceTypeEnum {
	return ScalingV2PolicyDetailScalingResourceTypeEnum{
		SCALING_GROUP: ScalingV2PolicyDetailScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: ScalingV2PolicyDetailScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c ScalingV2PolicyDetailScalingResourceType) Value() string {
	return c.value
}

func (c ScalingV2PolicyDetailScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingV2PolicyDetailScalingResourceType) UnmarshalJSON(b []byte) error {
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

type ScalingV2PolicyDetailPolicyStatus struct {
	value string
}

type ScalingV2PolicyDetailPolicyStatusEnum struct {
	INSERVICE ScalingV2PolicyDetailPolicyStatus
	PAUSED    ScalingV2PolicyDetailPolicyStatus
	EXECUTING ScalingV2PolicyDetailPolicyStatus
}

func GetScalingV2PolicyDetailPolicyStatusEnum() ScalingV2PolicyDetailPolicyStatusEnum {
	return ScalingV2PolicyDetailPolicyStatusEnum{
		INSERVICE: ScalingV2PolicyDetailPolicyStatus{
			value: "INSERVICE",
		},
		PAUSED: ScalingV2PolicyDetailPolicyStatus{
			value: "PAUSED",
		},
		EXECUTING: ScalingV2PolicyDetailPolicyStatus{
			value: "EXECUTING",
		},
	}
}

func (c ScalingV2PolicyDetailPolicyStatus) Value() string {
	return c.value
}

func (c ScalingV2PolicyDetailPolicyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingV2PolicyDetailPolicyStatus) UnmarshalJSON(b []byte) error {
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

type ScalingV2PolicyDetailScalingPolicyType struct {
	value string
}

type ScalingV2PolicyDetailScalingPolicyTypeEnum struct {
	ALARM      ScalingV2PolicyDetailScalingPolicyType
	SCHEDULED  ScalingV2PolicyDetailScalingPolicyType
	RECURRENCE ScalingV2PolicyDetailScalingPolicyType
}

func GetScalingV2PolicyDetailScalingPolicyTypeEnum() ScalingV2PolicyDetailScalingPolicyTypeEnum {
	return ScalingV2PolicyDetailScalingPolicyTypeEnum{
		ALARM: ScalingV2PolicyDetailScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ScalingV2PolicyDetailScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ScalingV2PolicyDetailScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ScalingV2PolicyDetailScalingPolicyType) Value() string {
	return c.value
}

func (c ScalingV2PolicyDetailScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingV2PolicyDetailScalingPolicyType) UnmarshalJSON(b []byte) error {
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
